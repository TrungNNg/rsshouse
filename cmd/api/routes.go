package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) route() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", app.test)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	return app.logRequest(app.recoverPanic(app.rateLimit(router)))
}
