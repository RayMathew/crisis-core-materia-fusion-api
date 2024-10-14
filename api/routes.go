package main

import (
	"net/http"

	_ "github.com/RayMathew/crisis-core-materia-fusion-api/api/docs"
	"github.com/julienschmidt/httprouter"

	httpSwagger "github.com/swaggo/http-swagger"
)

func (app *application) routes() http.Handler {
	mux := httprouter.New()

	mux.NotFound = http.HandlerFunc(app.notFound)
	mux.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowed)

	// Serve the Swagger UI
	mux.Handler("GET", "/docs/*any", httpSwagger.WrapHandler)

	mux.HandlerFunc("GET", "/status", app.status)
	mux.HandlerFunc("GET", "/materia", app.getAllMateria)

	// Adding content-type check middleware to only the POST method
	mux.Handler("POST", "/fusion", app.contentTypeCheck(http.HandlerFunc(app.fuseMateria)))

	return app.chainMiddlewares(mux)
}

func (app *application) chainMiddlewares(next http.Handler) http.Handler {
	middlewares := []func(http.Handler) http.Handler{
		app.recoverPanic,
		app.apiTimeout,
		app.rateLimiter,
	}

	for _, middleware := range middlewares {
		next = middleware(next)
	}

	return next
}
