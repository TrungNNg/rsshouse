package jobs

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/TrungNNg/rsshouse/internal/api"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/google/uuid"
	"github.com/mmcdole/gofeed"
)

// update feeds periodically
// 1. fetch all feeds that have last_fetched_at more than 1 hour
// 2. for each feed, run a goroutine to fetch and update the feed
func Aggregate(c *api.ApiConfig) {
	ticker := time.NewTicker(c.FetchFeedInterval)
	log.Printf("Aggregate job run every %s\n", c.FetchFeedInterval)

	for {
		<-ticker.C
		log.Printf("Begin updating feeds at %s\n", time.Now().String())
		feeds, err := c.DB.GetFeedsToFetch(context.Background(), sql.NullTime{Time: time.Now().UTC()})
		if err != nil {
			log.Printf("Error get feeds to update: %v\n", err)
		}
		for _, feed := range feeds {
			scrapeFeed(c, feed)
		}
	}
}

func scrapeFeed(c *api.ApiConfig, dbfeed database.Feed) {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	// fetch feed using gofetch
	feed, err := c.Parser.ParseURLWithContext(dbfeed.FeedLink, ctx)
	if err != nil {
		log.Printf("Error fetch feed: %v\n", err)
		return
	}

	// feed.Image can be nil
	var feedImgURL, feedImgTitle string
	if feed.Image != nil {
		feedImgTitle, feedImgURL = feed.Image.Title, feed.Image.URL
	}

	// feed.PublishedTime and feed.UpdatedTime can be nil
	var feedPublishedTime, feedUpdatedTime sql.NullTime
	if feed.PublishedParsed != nil {
		feedPublishedTime = sql.NullTime{Time: *feed.PublishedParsed, Valid: true}
	}
	if feed.UpdatedParsed != nil {
		feedUpdatedTime = sql.NullTime{Time: *feed.UpdatedParsed, Valid: true}
	}

	// update feed
	c.DB.UpdateFeedByID(ctx, database.UpdateFeedByIDParams{
		ID:              dbfeed.ID,
		UpdatedAt:       time.Now().UTC(),
		Title:           feed.Title,
		Descrip:         feed.Description,
		Link:            feed.Link,
		FeedLink:        feed.FeedLink,
		UpdatedParsed:   feedUpdatedTime,
		PublishedParsed: feedPublishedTime,
		Lang:            feed.Language,
		ImgUrl:          feedImgURL,
		ImgTitle:        feedImgTitle,
		FeedType:        feed.FeedType,
		LastFetchedAt:   sql.NullTime{Time: time.Now().UTC(), Valid: true},
	})

	// when update posts, need to remove old post in db as well.
	// we can just remove all then add all, but that might result in unnecessary db ops
	// when the feed doesn't have any new post.
	// It's better to decide whether to add or delete.
	// 1. fetch new post from feed (already done) -> []*gofeed.Item -> convert to []database.Post
	// 2. fetch all old post from feed -> []database.Post
	// 3. Turn them to set (aka map)
	// 4. compare to get a list of new posts to add -> new set - old set
	// 5. compare to get a list of post to remove -> old set - new set
	// 6. delete old post, update exiting post, add new post

	oldPosts, err := c.DB.GetPostsOfFeed(ctx, dbfeed.ID)
	if err != nil {
		log.Printf("Error fetch old posts for feed: %v\n", err)
		return
	}

	// newPostset
	newPostsMap := make(map[string]*gofeed.Item)
	for _, item := range feed.Items {
		newPostsMap[item.GUID] = item
	}

	// oldPostSet
	oldPostsMap := make(map[string]database.Post)
	for _, post := range oldPosts {
		oldPostsMap[post.Guid] = post
	}

	// Identify posts to add and to update
	postsToAdd := []*gofeed.Item{}
	postsToUpdate := []*gofeed.Item{}
	for guid, item := range newPostsMap {
		if _, exists := oldPostsMap[guid]; !exists {
			postsToAdd = append(postsToAdd, item)
		}
		if _, exists := oldPostsMap[guid]; exists {
			postsToUpdate = append(postsToUpdate, item)
		}
	}

	// Identify posts to remove
	postsToRemove := []database.Post{}
	for guid, post := range oldPostsMap {
		if _, exists := newPostsMap[guid]; !exists {
			postsToRemove = append(postsToRemove, post)
		}
	}

	// delete old posts
	for _, post := range postsToRemove {
		err = c.DB.DeletePostByID(ctx, post.ID)
		if err != nil {
			log.Printf("Error delete old post by id: %v\n", err)
		}
	}

	// update posts
	for _, post := range postsToUpdate {
		// post.Image can be nil
		var postImgURL, postImgTitle string
		if post.Image != nil {
			postImgTitle, postImgURL = post.Image.Title, post.Image.URL
		}

		// post.UpdatedParsed and post.PublishedParsed can be nil
		var postUpdateTime, postPublishedTime sql.NullTime
		if post.UpdatedParsed != nil {
			postUpdateTime = sql.NullTime{Time: *post.UpdatedParsed, Valid: true}
		}
		if post.PublishedParsed != nil {
			postPublishedTime = sql.NullTime{Time: *post.PublishedParsed, Valid: true}
		}

		err = c.DB.UpdatePostByGuid(ctx, database.UpdatePostByGuidParams{
			UpdatedAt:       time.Now().UTC(),
			Title:           post.Title,
			Descrip:         post.Description,
			PostLink:        post.Link,
			UpdatedParsed:   postUpdateTime,
			PublishedParsed: postPublishedTime,
			ImgUrl:          postImgURL,
			ImgTitle:        postImgTitle,
			FeedID:          dbfeed.ID,
			Guid:            post.GUID,
		})
		if err != nil {
			log.Println("error updating old post: ", err)
		}
	}

	// add new posts
	for _, post := range postsToAdd {
		// post.Image can be nil
		var postImgURL, postImgTitle string
		if post.Image != nil {
			postImgTitle, postImgURL = post.Image.Title, post.Image.URL
		}

		// post.UpdatedParsed and post.PublishedParsed can be nil
		var postUpdateTime, postPublishedTime sql.NullTime
		if post.UpdatedParsed != nil {
			postUpdateTime = sql.NullTime{Time: *post.UpdatedParsed, Valid: true}
		}
		if post.PublishedParsed != nil {
			postPublishedTime = sql.NullTime{Time: *post.PublishedParsed, Valid: true}
		}

		err = c.DB.AddPost(ctx, database.AddPostParams{
			ID:              uuid.New(),
			CreatedAt:       time.Now().UTC(),
			UpdatedAt:       time.Now().UTC(),
			Title:           post.Title,
			Descrip:         post.Description,
			PostLink:        post.Link,
			UpdatedParsed:   postUpdateTime,
			PublishedParsed: postPublishedTime,
			ImgUrl:          postImgURL,
			ImgTitle:        postImgTitle,
			Guid:            post.GUID,
			FeedID:          dbfeed.ID,
		})
		if err != nil {
			log.Println("error adding new post: ", err)
		}
	}
	log.Printf("finish scrape feed %s\n", dbfeed.Title)
}
