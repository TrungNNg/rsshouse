package api

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/google/uuid"
)

func (c *ApiConfig) SavePost(w http.ResponseWriter, r *http.Request) {
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
		Title    string `json:"title"`
		PostLink string `json:"post_link"`
	}{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.Decode(&reqData)

	// Check if a post is already in saved_posts by querying
	// for a record with a matching post_link.
	savedPost, err := c.DB.GetSavedPostByPostLink(r.Context(), reqData.PostLink)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		respondWithError(w, http.StatusInternalServerError, "Couldn't get saved post", err)
		return
	}

	// save post if not saved yet
	if errors.Is(err, sql.ErrNoRows) {
		// created new saved post in saved_posts table
		savedPost, err = c.DB.AddSavedPost(r.Context(), database.AddSavedPostParams{
			ID:        uuid.New(),
			CreatedAt: time.Now().UTC(),
			UpdatedAt: time.Now().UTC(),
			Title:     reqData.Title,
			PostLink:  reqData.PostLink,
		})
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't create save post", err)
			return
		}
	}

	// save post for user
	err = c.DB.UserSavePost(r.Context(), database.UserSavePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now().UTC(),
		UpdatedAt:   time.Now().UTC(),
		SavedAt:     time.Now().UTC(),
		UserID:      userID,
		SavedPostID: savedPost.ID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't save post", err)
		return
	}
	respondWithJSON(w, http.StatusOK, "Saved post :)")
}
