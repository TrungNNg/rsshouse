package api

import (
	"net/http"

	"github.com/TrungNNg/rsshouse/internal/auth"
)

func (c *ApiConfig) Test(w http.ResponseWriter, r *http.Request) {
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

	respondWithJSON(w, http.StatusOK, userID.String())
}
