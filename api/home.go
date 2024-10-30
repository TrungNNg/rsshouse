package api

import (
	"fmt"
	"net/http"
)

func (c *ApiConfig) HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to RSSHouse!")
}
