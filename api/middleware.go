package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/didip/tollbooth"
)

func (app *application) recoverPanic(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()
			if err != nil {
				app.serverError(w, r, fmt.Errorf("%s", err))
			}
		}()

		next.ServeHTTP(w, r)
	})
}

func (app *application) contentTypeCheck(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Content-Type") != "application/json" {
			app.unsupportedMediaType(w, r)

			return
		}
		next.ServeHTTP(w, r)
	})
}

func (app *application) rateLimiter(next http.Handler) http.Handler {
	limiter := tollbooth.NewLimiter(1, nil)
	limiter.SetIPLookups([]string{"X-Real-IP", "X-Forwarded-For", "RemoteAddr"})

	return tollbooth.LimitHandler(limiter, next)
}

func (app *application) apiTimeout(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		timeoutDuration := time.Duration(app.config.apiTimeout) * time.Second
		// Create a context with the specified timeout
		ctx, cancel := context.WithTimeout(r.Context(), timeoutDuration)
		defer cancel() // Make sure to cancel the context to free resources

		// Update the request with the new context
		r = r.WithContext(ctx)

		// Create a channel to signal completion
		done := make(chan struct{})

		// Run the next handler in a goroutine
		go func() {
			next.ServeHTTP(w, r)
			close(done)
		}()

		select {
		case <-done:
			// Handler completed successfully
			return
		case <-ctx.Done():
			// Timeout occurred
			app.gatewayTimeout(w, r)
			return
		}
	})
}
