package api

import (
	"encoding/json"
	"net/http"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/google/uuid"
)

func (c *ApiConfig) UnsavePost(w http.ResponseWriter, r *http.Request) {
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
		PostID string `json:"post_id"`
	}{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	decoder.Decode(&reqData)

	postID, err := uuid.Parse(reqData.PostID)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid post ID", err)
		return
	}

	// unsave post by remove row from user_saved_posts join table
	err = c.DB.UnsavePost(r.Context(), database.UnsavePostParams{
		UserID:      userID,
		SavedPostID: postID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't unsave post", err)
		return
	}

	// check if post is still saved by any other users, if not delete post from saved table
	remain, err := c.DB.CountSaved(r.Context(), postID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Something went wrong with unsave post", err)
		return
	}

	if remain == 0 {
		err = c.DB.DeleteSavedPost(r.Context(), postID)
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Something went wrong with unsave post", err)
			return
		}
	}

	respondWithJSON(w, http.StatusOK, "Unsaved post")
}
