package main

import (
	"net/http"

	"github.com/MadMorris/wms-base-rest-api/src/api/handlers"
)

func main() {
	// Create a new router
	http.HandleFunc("/user", handlers.CreateUserHandler)

	// Start server
	http.ListenAndServe(":8080", nil)
}
