package api

import (
	"encoding/json"
	"net/http"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
)

func (c *ApiConfig) UnsubcribeFeed(w http.ResponseWriter, r *http.Request) {
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
		FeedLink string `json:"feed_link"`
	}{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.Decode(&reqData)

	dbFeed, err := c.DB.GetFeedByFeedLink(r.Context(), reqData.FeedLink)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something when wrong", err)
		return
	}

	err = c.DB.UnsubcribeFeed(r.Context(), database.UnsubcribeFeedParams{
		UserID: userID,
		FeedID: dbFeed.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something when wrong", err)
		return
	}
	respondWithJSON(w, http.StatusOK, "User unsubcribed to feed")
}
