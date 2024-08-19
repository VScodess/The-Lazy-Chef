package main

import (
	"The-Lazy-Chef/backend/database"
	"The-Lazy-Chef/backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	// Connect to MongoDB
	database.Connect()

	// Set up router with routes
	router := mux.NewRouter()
	routes.SetupRoutes(router)

	log.Println("Backend server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
