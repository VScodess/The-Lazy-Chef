package main

import (
	"The-Lazy-Chef/backend/config"
	"The-Lazy-Chef/backend/database"
	"The-Lazy-Chef/backend/routes"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {

	// Connect to MongoDB
	database.Connect()

	// Load config
	cfg := config.LoadConfig()

	// Set up router with routes
	router := mux.NewRouter()
	routes.SetupRoutes(router, cfg)

	// Enable CORS for all origins
	corsOptions := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	log.Println("Backend server running on port 8080")
	log.Fatal(http.ListenAndServe(":8080", corsOptions(router)))
}
