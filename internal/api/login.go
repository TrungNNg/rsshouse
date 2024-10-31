package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
)

// send user access token and refresh token on success login
// new refresh token is save in db
func (c *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Can not decode json", err)
		return
	}

	// find user with given username
	dbUser, err := c.DB.GetUserByUsername(r.Context(), reqData.Username)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid username or password", err)
		return
	}

	// compare the given password with the hashed in db
	err = auth.CheckPasswordHash(reqData.Password, dbUser.HashedPassword)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid username or password", err)
		return
	}

	// create new jwt token and send to client
	jwt, err := auth.MakeJWT(dbUser.ID, c.Secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create token", err)
		return
	}

	// create new refresh token in db
	refreshToken, err := auth.MakeRefreshToken()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create refresh token", err)
		return
	}

	_, err = c.DB.CreateRefreshToken(r.Context(), database.CreateRefreshTokenParams{
		Token:     refreshToken,
		UserID:    dbUser.ID,
		ExpiresAt: time.Now().UTC().Add(time.Hour * 24 * 60),
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't save refresh token", err)
		return
	}

	// send access token and refresh token to client
	resData := struct {
		AccessToken  string `json:"token"`
		RefreshToken string `json:"refresh_token"`
	}{
		AccessToken:  jwt,
		RefreshToken: refreshToken,
	}

	respondWithJSON(w, http.StatusOK, resData)
}
