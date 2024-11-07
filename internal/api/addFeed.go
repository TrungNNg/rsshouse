package api

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"strings"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/google/uuid"
)

// authenticated endpoint, take in json: {"url":"string"}
func (c *ApiConfig) AddFeed(w http.ResponseWriter, r *http.Request) {
	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find JWT", err)
		return
	}

	userID, err := auth.ValidateJWT(token, c.Secret)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid access token", err)
		return
	}

	// TODO: when add a feed user also need to pick at least 1 topic and at most
	// 5 topics that the feed is about. There will be a simple check in front end
	// but that can not be rely on, so there will be a check in backend as well

	// count topic here

	reqData := struct {
		URL string `json:"url"`
	}{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.Decode(&reqData)

	feed, err := c.Parser.ParseURL(reqData.URL)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid feed url", err)
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

	// save feed to db
	dbFeed, err := c.DB.AddFeed(r.Context(), database.AddFeedParams{
		ID:              uuid.New(),
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
		UserID:          userID,
	})
	if err != nil {
		if strings.Contains(err.Error(), `duplicate key value violates unique constraint "feeds_feed_link_key"`) {
			respondWithError(w, http.StatusBadRequest, "This url already added :)", err)
			return
		}
		respondWithError(w, http.StatusBadRequest, "Coudn't parse feed url", err)
		return
	}

	// save post to db
	for _, post := range feed.Items {
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

		err = c.DB.AddPost(r.Context(), database.AddPostParams{
			ID:              uuid.New(),
			Title:           post.Title,
			Descrip:         post.Description,
			PostLink:        post.Link,
			UpdatedParsed:   postUpdateTime,
			PublishedParsed: postPublishedTime,
			ImgUrl:          postImgURL,
			ImgTitle:        postImgTitle,
			Guid:            post.GUID,
			FeedID:          dbFeed.ID,
		})
		log.Println("error adding post: ", err)
	}
	respondWithJSON(w, http.StatusOK, "Add feed successfuly")
}
