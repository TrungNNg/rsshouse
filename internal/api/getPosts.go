package api

import (
	"net/http"

	"github.com/TrungNNg/rsshouse/internal/auth"
)

func (c *ApiConfig) GetPosts(w http.ResponseWriter, r *http.Request) {
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

	posts, err := c.DB.GetSubcribedPostsOfUser(r.Context(), userID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Coudn't get posts from subcribed feeds", err)
		return
	}

	if len(posts) == 0 {
		respondWithJSON(w, http.StatusOK, "No posts found")
		return
	}

	for _, p := range posts {
		println(p.Guid)
	}
}
