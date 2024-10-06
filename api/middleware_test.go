package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecoverPanic(t *testing.T) {
	app := newTestApplication()

	// Create a handler that panics
	panicHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		panic("test panic")
	})

	// Wrap the panicHandler with the recoverPanic middleware
	handler := app.recoverPanic(panicHandler)

	// Create a test request and response recorder
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check if a server error is written after panic
	assert.Equal(t, http.StatusInternalServerError, rr.Code)
}

func TestContentTypeCheck(t *testing.T) {
	app := newTestApplication()

	// Create a test handler that does nothing
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Wrap with contentTypeCheck middleware
	handler := app.contentTypeCheck(testHandler)

	// Test: Correct Content-Type
	req := httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Test: Incorrect Content-Type
	req = httptest.NewRequest(http.MethodPost, "/", nil)
	req.Header.Set("Content-Type", "text/plain")
	rr = httptest.NewRecorder()

	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusUnsupportedMediaType, rr.Code)
}

func TestRateLimiter(t *testing.T) {
	app := newTestApplication()

	// Create a test handler that does nothing
	testHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	// Wrap with rateLimiter middleware
	handler := app.rateLimiter(testHandler)

	// Test: First request should succeed
	req := httptest.NewRequest(http.MethodGet, "/", nil)
	rr := httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusOK, rr.Code)

	// Test: Second request should be rate-limited
	rr = httptest.NewRecorder()
	handler.ServeHTTP(rr, req)
	assert.Equal(t, http.StatusTooManyRequests, rr.Code)
}

// getting a race condition

// func TestApiTimeout(t *testing.T) {
// 	app := newTestApplication()

// 	// Set a short timeout for testing
// 	app.config.apiTimeout = 1 // 1 second timeout

// 	// Create a handler that simulates a long-running process
// 	longRunningHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		time.Sleep(2 * time.Second) // Simulate a delay
// 		w.WriteHeader(http.StatusOK)
// 	})

// 	// Wrap with apiTimeout middleware
// 	handler := app.apiTimeout(longRunningHandler)

// 	// Create a test request and response recorder
// 	req := httptest.NewRequest(http.MethodGet, "/", nil)
// 	rr := httptest.NewRecorder()

// 	// Call the handler directly (no goroutine needed)
// 	handler.ServeHTTP(rr, req)

// 	// Check if the request timed out
// 	assert.Equal(t, http.StatusGatewayTimeout, rr.Code)
// }
