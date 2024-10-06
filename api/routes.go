package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	mux.HandlerFunc("GET", "/status", app.status)
	mux.HandlerFunc("GET", "/materia", app.getAllMateria)
	mux.HandlerFunc("POST", "/fusion", app.fuseMateria)

	return app.chainMiddlewares(mux)
}

func (app *application) chainMiddlewares(next http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{
		app.contentTypeMiddleware,
		app.recoverPanic,
		app.rateLimiter,
	}

	for _, middleware := range middlewares {
		next = middleware(next)
	}

	return next
}
