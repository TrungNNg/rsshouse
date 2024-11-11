package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
	"github.com/google/uuid"
)

func (c *ApiConfig) SignUp(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	err := decoder.Decode(&reqData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode json", err)
		return
	}

	err = checkUsernameAndPassword(reqData.Username, reqData.Password)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error(), nil)
		return
	}

	hashedPassword, err := auth.HashPassword(reqData.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Problem with password", err)
		return
	}

	dbUser, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:             uuid.New(),
		CreatedAt:      time.Now().UTC(),
		UpdatedAt:      time.Now().UTC(),
		Username:       reqData.Username,
		HashedPassword: hashedPassword,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't create user", err)
		return
	}

	// For debuging
	fmt.Println("Create user successful")
	fmt.Println(User{
		ID:             dbUser.ID,
		CreatedAt:      dbUser.CreatedAt,
		UpdatedAt:      dbUser.UpdatedAt,
		Username:       dbUser.Username,
		HashedPassword: dbUser.HashedPassword,
	}) //

	respondWithJSON(w, http.StatusCreated, "User created")
}
