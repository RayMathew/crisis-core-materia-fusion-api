package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"strings"

	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/response"
	"github.com/RayMathew/crisis-core-materia-fusion-api/internal/validator"
)

func (app *application) reportServerError(r *http.Request, err error) {
	var (
		message = err.Error()
		method  = r.Method
		url     = r.URL.String()
		trace   = string(debug.Stack())
	)

	requestAttrs := slog.Group("request", "method", method, "url", url)
	app.logger.Error(message, requestAttrs, "trace", trace)
}

func (app *application) errorMessage(w http.ResponseWriter, r *http.Request, status int, message string, headers http.Header) {
	// Exit early if a response has already been written
	if w.Header().Get("Content-Type") != "" {
		return
	}

	message = strings.ToUpper(message[:1]) + message[1:]

	err := response.JSONWithHeaders(w, status, ErrorResponseDTO{Error: message}, headers)
	if err != nil {
		app.reportServerError(r, err)
		if w.Header().Get("Content-Type") == "" {
			w.WriteHeader(http.StatusInternalServerError)
		}
	}
}

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.reportServerError(r, err)

	message := "The server encountered a problem and could not process your request"
	app.errorMessage(w, r, http.StatusInternalServerError, message, nil)
}

func (app *application) notFound(w http.ResponseWriter, r *http.Request) {
	message := "The requested resource could not be found"
	app.errorMessage(w, r, http.StatusNotFound, message, nil)
}

func (app *application) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s method is not supported for this resource", r.Method)
	app.errorMessage(w, r, http.StatusMethodNotAllowed, message, nil)
}

func (app *application) unsupportedMediaType(w http.ResponseWriter, r *http.Request) {
	message := fmt.Sprintf("The %s Content-Type is not supported", r.Header.Get("Content-Type"))
	app.errorMessage(w, r, http.StatusUnsupportedMediaType, message, nil)
}

func (app *application) gatewayTimeout(w http.ResponseWriter, r *http.Request) {
	message := "Request timed out"
	app.errorMessage(w, r, http.StatusGatewayTimeout, message, nil)
}

func (app *application) badRequest(w http.ResponseWriter, r *http.Request, err error) {
	app.errorMessage(w, r, http.StatusBadRequest, err.Error(), nil)
}

func (app *application) failedValidation(w http.ResponseWriter, r *http.Request, v validator.Validator) {
	err := response.JSON(w, http.StatusUnprocessableEntity, v)
	if err != nil {
		app.serverError(w, r, err)
	}
}
