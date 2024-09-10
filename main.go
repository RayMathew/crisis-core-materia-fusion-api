package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

func main() {

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)

	listenAddr := "localhost:5001"

	server := &http.Server{
		Addr:     listenAddr,
		Handler:  routes(), //tracing(nextRequestID)(logging(logger)(routes())),
		ErrorLog: logger,
		// ReadTimeout:  5 * time.Second,
		// WriteTimeout: 10 * time.Second,
		// IdleTimeout:  15 * time.Second,
	}

	// Listen for CTRL+C or kill and start shutting down the app without
	// disconnecting people by not taking any new requests. ("Graceful Shutdown")
	// done := make(chan bool)
	// quit := make(chan os.Signal, 1)
	// signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	// go func() {
	// 	<-quit
	// 	logger.Println("Server is shutting down...")
	// 	atomic.StoreInt32(&healthy, 0)

	// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	// 	defer cancel()

	// 	server.SetKeepAlivesEnabled(false)
	// 	if err := server.Shutdown(ctx); err != nil {
	// 		logger.Fatalf("Could not gracefully shutdown the server: %v\n", err)
	// 	}
	// 	close(done)
	// }()

	logger.Println("Server is ready to handle requests at", listenAddr)
	// atomic.StoreInt32(&healthy, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("Could not listen on %s: %v\n", listenAddr, err)
	}

	// <-done
	logger.Println("Server stopped")
}

func routes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/", indexHandler)
	// router.HandleFunc("/health", healthHandler)
	// router.HandleFunc("/hello", helloHandler)
	// router.HandleFunc("/json-as-text", forceTextHandler)
	return router
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode("hello world")
}
