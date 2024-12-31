package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// let implement rate limiter
// what is a rate limiter, basically i don't want a user to keep hitting an endpoint too many
// time in too short period of time, which will make server create lot of go routine to handle the
// request.

// one simple way to implement rate limiter is to keep track of an ip adress and
// previous hitting time. I can use a map to do this

// 1. how to get ip address from request? there is probably code in std to handle this
// where to save the map that keep track of ip and time ? -> global var in the server ?
// is there a better place ?
// why middleware? what is middleware ??????
// what is the time zone of the time, i guess it make sense to use the server time
// but what if there is numtiple server on diffrent computer in different time zone?
//

func (app *application) route() http.Handler {
	router := httprouter.New()
	router.HandlerFunc(http.MethodGet, "/", app.test)

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)

	return app.logRequest(app.recoverPanic(app.rateLimit(router)))
}
