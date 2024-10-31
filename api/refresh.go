package api

import (
	"net/http"
	"time"

	"github.com/TrungNNg/rsshouse/internal/auth"
)

// get new access token using refresh token, if expired return 401
func (c *ApiConfig) Refresh(w http.ResponseWriter, r *http.Request) {
	refreshToken, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't find refresh token", err)
		return
	}

	dbUser, err := c.DB.GetUserFromRefreshToken(r.Context(), refreshToken)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Couldn't get user for refresh token", err)
		return
	}

	// use refresh token to generate new access token which store user's id
	jwtToken, err := auth.MakeJWT(dbUser.ID, c.Secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create access token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{
		Token: jwtToken,
	})
}
