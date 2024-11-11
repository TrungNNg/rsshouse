package jobs

import (
	"context"
	"database/sql"
	"log"
	"time"

	"github.com/TrungNNg/rsshouse/internal/api"
	"github.com/TrungNNg/rsshouse/internal/database"
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

	// when update posts, need to remove old post as well
	// we can just remove all then add all, but that might result in unnecessary db ops
	// when the feed doesn't have any new post.
	// 1. fetch new post from feed
	// 2. fetch all old post from feed
	// 3. compare to get a list of new posts to add
	// 4. compare to get a list of post to remove
	// 5.
}
