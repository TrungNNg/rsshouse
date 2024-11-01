package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

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

	// save feed to db
	_, err = c.DB.AddFeed(r.Context(), database.AddFeedParams{
		ID:            uuid.New(),
		Title:         feed.Title,
		Descrip:       feed.Description,
		FeedLink:      feed.FeedLink,
		UpdatedParsed: time.Now().UTC(),
		Lang:          feed.Language,
		ImgUrl:        feed.Image.URL,
		ImgTitle:      feed.Image.Title,
		FeedType:      feed.FeedType,
		UserID:        userID,
	})
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint \"feeds_feed_link_key\"") {
			respondWithError(w, http.StatusBadRequest, "This url already added :)", err)
			return
		}
		respondWithError(w, http.StatusBadRequest, "Coudn't parse feed url", err)
		return
	}

	// save post to db
	for _, post := range feed.Items {
		fmt.Println(post.Title)
	}
}
