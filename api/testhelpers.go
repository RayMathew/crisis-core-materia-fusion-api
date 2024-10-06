// testhelpers.go (or similar)

package main

import (
	"log/slog"
	"os"
	"sync"
)

// newTestApplication creates a mock application for unit tests
func newTestApplication() *application {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))

	cfg := config{
		baseURL:                  "http://localhost:4000",
		httpPort:                 4000,
		apiTimeout:               5,
		apiCallsAllowedPerSecond: 3,
	}

	return &application{
		db:     nil,
		logger: logger,
		cache:  make(map[string]interface{}),
		wg:     sync.WaitGroup{},
		mu:     sync.Mutex{},
		config: cfg,
	}
}
