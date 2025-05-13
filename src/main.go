package main

import (
	"log"
	"net/http"
)

func main() {
	// Serve files from the "static" directory.
	// This will handle all requests and serve static files relative to the "./static" path.
	fileServer := http.FileServer(http.Dir("./static"))

	// Handle root URL "/" by serving files using the fileServer handler.
	http.Handle("/", fileServer)

	// Define the port the server will listen on.
	port := ":8080"
	log.Printf("Starting server on http://localhost%s\n", port)

	// Start the HTTP server and log any fatal errors.
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
