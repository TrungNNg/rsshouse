package api

import (
	"encoding/json"
	"log"
	"net/http"
)

// send the msg to client, log out err in terminal
func respondWithError(w http.ResponseWriter, code int, msg string, err error) {
	if err != nil {
		log.Println(err)
	}
	if code > 499 {
		log.Printf("Responding with 5XX error: %s", msg)
	}
	type errorResponse struct {
		Error string `json:"error"`
	}
	respondWithJSON(w, code, errorResponse{
		Error: msg,
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func CheckUsernameAndPassword(username, password string) error {
	// Check username length
	if len(username) < 2 {
		return ErrUsernameTooShort
	}
	if len(username) > 30 {
		return ErrUsernameTooLong
	}

	// Check password length
	if len(password) < 5 {
		return ErrPasswordTooShort
	}
	if len(password) > 30 {
		return ErrPasswordTooLong
	}

	return nil
}
