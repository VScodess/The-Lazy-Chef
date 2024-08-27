package handlers

import (
	"The-Lazy-Chef/backend/config"
	"net/http"

	"github.com/gorilla/mux"
)

// GetRecipes retrieves all recipes from the database.
func GetRecipes(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	recipes, err := fetchRecipes(cfg)
	if err != nil {
		http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, recipes)
}

// CreateRecipe adds a new recipe to the database.
func CreateRecipe(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	recipe, err := parseRecipeForm(r)
	if err != nil {
		http.Error(w, "Failed to parse recipe", http.StatusBadRequest)
		return
	}

	if err := insertRecipe(cfg, recipe); err != nil {
		http.Error(w, "Failed to save recipe", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, recipe)
}

// GetRecipe retrieves a single recipe by ID.
func GetRecipe(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	response, err := fetchRecipeByID(cfg, id)
	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}

// GetRecipesByCategory retrieves recipes by their category.
func GetRecipesByCategory(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	category := mux.Vars(r)["category"]

	// Fetch recipes by category
	recipes, err := fetchRecipesByCategory(cfg, category)
	if err != nil {
		http.Error(w, "Failed to fetch recipes", http.StatusInternalServerError)
		return
	}

	if len(recipes) == 0 {
		http.Error(w, "No recipes found for the given category", http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, recipes)
}

// UpdateRecipe updates an existing recipe by ID.
func UpdateRecipe(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	recipe, err := parseRecipeForm(r)
	if err != nil {
		http.Error(w, "Failed to parse recipe", http.StatusBadRequest)
		return
	}

	if err := updateRecipeByID(cfg, id, recipe); err != nil {
		http.Error(w, "Failed to update recipe", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, recipe)
}

// DeleteRecipe deletes a recipe by ID.
func DeleteRecipe(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := deleteRecipeByID(cfg, id); err != nil {
		http.Error(w, "Failed to delete recipe", http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Recipe deleted successfully"})
}
