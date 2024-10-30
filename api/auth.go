package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TrungNNg/rsshouse/internal/auth"
	"github.com/TrungNNg/rsshouse/internal/database"
)

func (c *ApiConfig) SignUp(w http.ResponseWriter, r *http.Request) {
	reqData := struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}{}

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&reqData)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Can not decode json", err)
		return
	}

	if len(reqData.Password) < 5 {
		respondWithError(w, http.StatusBadRequest, "Password too short", nil)
		return
	}

	hashedPassword, err := auth.HashPassword(reqData.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Problem with password", err)
		return
	}

	dbUser, err := c.DB.CreateUser(r.Context(), database.CreateUserParams{
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
	})
	respondWithJSON(w, http.StatusCreated, "User created")
}

func (c *ApiConfig) Login(w http.ResponseWriter, r *http.Request) {

}

func (c *ApiConfig) Logout(w http.ResponseWriter, r *http.Request) {

}
