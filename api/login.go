package api

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/TrungNNg/rsshouse/internal/auth"
)

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

	// create jwt token and send to client
	jwt, err := auth.MakeJWT(dbUser.ID, c.Secret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Can not create token", err)
		return
	}

	respondWithJSON(w, http.StatusOK, jwt)
}
