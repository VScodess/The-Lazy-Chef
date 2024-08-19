package routes

import (
	"The-Lazy-Chef/backend/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router) {

	// Get all recipes
	router.HandleFunc("/recipes", handlers.GetRecipes).Methods("GET")

	// Get all recipes by category
	router.HandleFunc("/recipes/{category}", handlers.GetRecipesByCategory).Methods("GET")

	// Create a recipe
	router.HandleFunc("/recipes", handlers.CreateRecipe).Methods("POST")

	// Get a recipe
	router.HandleFunc("/recipes/{id}", handlers.GetRecipe).Methods("GET")

	// Update a recipe
	router.HandleFunc("/recipes/{id}", handlers.UpdateRecipe).Methods("PUT")

	// Delete a recipe
	router.HandleFunc("/recipes/{id}", handlers.DeleteRecipe).Methods("DELETE")

}