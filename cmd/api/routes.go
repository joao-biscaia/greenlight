package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)
	standard := alice.New(app.recoverPanic, app.enableCors, app.rateLimit, app.authenticate)

	router.Handler(http.MethodGet, "/v1/healthcheck", standard.ThenFunc(app.healthcheckHandler))

	router.Handler(http.MethodGet, "/v1/movies", app.requirePermission("movies:read", standard.ThenFunc(app.listMoviesHandler)))
	router.Handler(http.MethodPost, "/v1/movies", app.requirePermission("movies:write", standard.ThenFunc(app.createMovieHandler)))
	router.Handler(http.MethodGet, "/v1/movies/:id", app.requirePermission("movies:read", standard.ThenFunc(app.showMovieHandler)))
	router.Handler(http.MethodPatch, "/v1/movies/:id", app.requirePermission("movies:write", standard.ThenFunc(app.updateMovieHandler)))
	router.Handler(http.MethodDelete, "/v1/movies/:id", app.requirePermission("movies:write", standard.ThenFunc(app.deleteMovieHandler)))

	router.Handler(http.MethodPost, "/v1/users", standard.ThenFunc(app.registerUserHandler))
	router.Handler(http.MethodPut, "/v1/users/activated", standard.ThenFunc(app.activateUserHandler))
	router.Handler(http.MethodPost, "/v1/tokens/authentication", standard.ThenFunc(app.createAuthenticationTokenHandler))

	router.Handler(http.MethodPost, "/v1/tokens/activation", standard.ThenFunc(app.createActivationTokenHandler))

	return standard.Then(router)
}
