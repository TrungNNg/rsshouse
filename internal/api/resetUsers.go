package api

import "net/http"

func (c *ApiConfig) ResetUsers(w http.ResponseWriter, r *http.Request) {
	err := c.DB.ResetUsers(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Can not reset users table", err)
		return
	}
}
