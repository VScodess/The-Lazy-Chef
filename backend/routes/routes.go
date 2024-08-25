package routes

import (
	"The-Lazy-Chef/backend/config"
	"The-Lazy-Chef/backend/handlers"

	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes(router *mux.Router, cfg *config.Config) {

	router.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetRecipes(cfg, w, r)
	}).Methods("GET")

	// Search recipes
	router.HandleFunc("/recipes/search", func(w http.ResponseWriter, r *http.Request) {
		handlers.SearchRecipes(cfg, w, r)
	}).Methods("GET")

	// Get all recipes by category
	router.HandleFunc("/recipes/{category}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetRecipesByCategory(cfg, w, r)
	}).Methods("GET")

	// Create a recipe
	router.HandleFunc("/recipes", func(w http.ResponseWriter, r *http.Request) {
		handlers.CreateRecipe(cfg, w, r)
	}).Methods("POST")

	// Get a recipe
	router.HandleFunc("/recipes/{category}/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.GetRecipe(cfg, w, r)
	}).Methods("GET")

	// Update a recipe
	router.HandleFunc("/recipes/{category}/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.UpdateRecipe(cfg, w, r)
	}).Methods("PUT")

	// Delete a recipe
	router.HandleFunc("/recipes/{category}/{id}", func(w http.ResponseWriter, r *http.Request) {
		handlers.DeleteRecipe(cfg, w, r)
	}).Methods("DELETE")
}
