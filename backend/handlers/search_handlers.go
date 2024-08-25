package handlers

import (
	"The-Lazy-Chef/backend/config"
	"encoding/json"
	"net/http"
)

// SearchRecipes allows users to search for recipes based on a query and category.
func SearchRecipes(cfg *config.Config, w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	category := r.URL.Query().Get("category")

	if query == "" || category == "" {
		http.Error(w, "Both search query and category are required", http.StatusBadRequest)
		return
	}

	recipes, err := searchRecipesByCategoryAndQuery(cfg, category, query)
	if err != nil {
		http.Error(w, "Failed to perform search", http.StatusInternalServerError)
		return
	}

	if len(recipes) == 0 {
		http.Error(w, "No recipes found matching the search criteria", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
