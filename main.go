package main

import (
	"log"
	"net/http"
	"os"

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

	log.Println("Listening on port " + os.Getenv("PORT"))
	err := http.ListenAndServe(":"+os.Getenv("PORT"),
		// Log all routes
		middleware.LogRequest(router.Mux),
	)

	if err != nil {
		panic(err)
	}
}
