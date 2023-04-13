package main

import (
	"fmt"
	"log"
	"net/http"
)

// loggingMiddleware is a middleware function that logs the URL path of the incoming request
func loggingMiddleware(next http.Handler) http.Handler {
	
	// Return a new http.Handler that executes the following code when a request is received
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Print the URL path of the incoming request
	
		// Call the next middleware or request handler in the chain
		next.ServeHTTP(w, r)
	})
}

// mainHandler is a request handler that writes a "Hello Go World!" message to the response writer
func mainHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello Go World!")
}

// mainHandlerapi1 is a request handler that writes an "api1" message to the response writer
func mainHandlerapi1(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "API1 hitted")
}

func main() {
	// Add the logging middleware to the mainHandler and map it to the root path "/"
	http.Handle("/", loggingMiddleware(http.HandlerFunc(mainHandler)))
	// Add the logging middleware to the mainHandlerapi1 and map it to the "/api1" path
	http.Handle("/api1", loggingMiddleware(http.HandlerFunc(mainHandlerapi1)))
	// Start the HTTP server on port 8000
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
