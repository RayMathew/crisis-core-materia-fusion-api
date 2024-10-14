// errors_test.go

package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServerError(t *testing.T) {
	app := newTestApplication() // Get the test app instance

	// Create a dummy HTTP request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Call the serverError function
	err := errors.New("test server error")
	app.serverError(rr, req, err)

	// Assert status code is 500
	require.Equal(t, http.StatusInternalServerError, rr.Code)

	// Assert the body content matches the expected message
	expected := `{"Error":"The server encountered a problem and could not process your request"}`
	require.JSONEq(t, expected, rr.Body.String())

	// Assert Content-Type header is correctly set to application/json
	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestNotFound(t *testing.T) {
	app := newTestApplication()

	req := httptest.NewRequest(http.MethodGet, "/nonexistent", nil)
	rr := httptest.NewRecorder()

	app.notFound(rr, req)

	require.Equal(t, http.StatusNotFound, rr.Code)

	expected := `{"Error":"The requested resource could not be found"}`
	require.JSONEq(t, expected, rr.Body.String())

	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestMethodNotAllowed(t *testing.T) {
	app := newTestApplication()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	app.methodNotAllowed(rr, req)

	require.Equal(t, http.StatusMethodNotAllowed, rr.Code)

	expected := `{"Error":"The POST method is not supported for this resource"}`
	require.JSONEq(t, expected, rr.Body.String())

	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestUnsupportedMediaType(t *testing.T) {
	app := newTestApplication()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/xml")
	rr := httptest.NewRecorder()

	app.unsupportedMediaType(rr, req)

	require.Equal(t, http.StatusUnsupportedMediaType, rr.Code)

	expected := `{"Error":"The application/xml Content-Type is not supported"}`
	require.JSONEq(t, expected, rr.Body.String())

	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestGatewayTimeout(t *testing.T) {
	app := newTestApplication()

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	app.gatewayTimeout(rr, req)

	require.Equal(t, http.StatusGatewayTimeout, rr.Code)

	expected := `{"Error":"Request timed out"}`
	require.JSONEq(t, expected, rr.Body.String())

	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestBadRequest(t *testing.T) {
	app := newTestApplication()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	err := errors.New("bad request error")

	app.badRequest(rr, req, err)

	require.Equal(t, http.StatusBadRequest, rr.Code)

	expected := `{"Error":"Bad request error"}`
	require.JSONEq(t, expected, rr.Body.String())

	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}

func TestFailedValidation(t *testing.T) {
	app := newTestApplication()

	req := httptest.NewRequest(http.MethodPost, "/", nil)
	rr := httptest.NewRecorder()

	var fusionReq MateriaFusionRequest
	fusionReq.Validator.AddFieldError("field", "must be provided")

	app.failedValidation(rr, req, fusionReq.Validator)

	require.Equal(t, http.StatusUnprocessableEntity, rr.Code)

	expected := `{"FieldErrors":{"field":"must be provided"}}`
	require.JSONEq(t, expected, rr.Body.String())

	require.Equal(t, "application/json", rr.Header().Get("Content-Type"))
}
