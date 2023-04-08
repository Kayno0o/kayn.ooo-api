package main

import (
	"net/http"

	middleware "kayn.ooo/api/src/Middleware"
	repository "kayn.ooo/api/src/Repository"
	router "kayn.ooo/api/src/Router"
)

func main() {
	// Connect to the database
	repository.Connect()

	// Initialize the server's routes
	router.InitRoutes()

	http.ListenAndServe(":8080",
		// Log all routes
		middleware.LogRequest(router.Mux),
	)
}
