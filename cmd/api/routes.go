package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) route() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", app.test)
	return router
}
