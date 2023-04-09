package main

import (
	"net/http"

	"github.com/joho/godotenv"
	middleware "kayn.ooo/api/src/Middleware"
	repository "kayn.ooo/api/src/Repository"
	router "kayn.ooo/api/src/Router"
)

func main() {
	godotenv.Load(".env")

	// Connect to the database
	repository.Connect()

	// Initialize the server's routes
	router.InitRoutes()

	//http.ListenAndServe and log errors
	err := http.ListenAndServe(":8080",
		// Log all routes
		middleware.LogRequest(router.Mux),
	)

	if err != nil {
		panic(err)
	}
}
