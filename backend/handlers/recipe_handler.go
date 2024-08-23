package handlers

import (
	"The-Lazy-Chef/backend/database"
	"The-Lazy-Chef/backend/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	var recipes []map[string]interface{}
	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// M{} when order doesnt matter and you want everything
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var recipe models.Recipe
		cursor.Decode(&recipe)

		responseRecipe := map[string]interface{}{
			"id":          recipe.ID,
			"name":        recipe.Name,
			"category":    recipe.Category,
			"ingredients": recipe.Ingredients,
			"steps":       recipe.Steps,
			"tags":        recipe.Tags,
			"summary":     recipe.Summary,
			"image":       base64.StdEncoding.EncodeToString(recipe.Image),
		}

		recipes = append(recipes, responseRecipe)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)

}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {

	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		http.Error(w, "Unable to parse form", http.StatusBadRequest)
		return
	}

	// Read the image file
	file, _, err := r.FormFile("image")
	if err != nil {
		http.Error(w, "Error retrieving the image file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Read the file into a byte slice
	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		http.Error(w, "Error reading the image file", http.StatusInternalServerError)
		return
	}

	// Extract other form fields
	name := r.FormValue("name")
	category := r.FormValue("category")
	ingredients := strings.Split(r.FormValue("ingredients"), ",")
	steps := strings.Split(r.FormValue("steps"), ",")
	tags := strings.Split(r.FormValue("tags"), ",")
	summary := r.FormValue("summary")

	recipe := models.Recipe{
		ID:          primitive.NewObjectID(),
		Name:        name,
		Category:    category,
		Ingredients: ingredients,
		Steps:       steps,
		Tags:        tags,
		Summary:     summary,
		Image:       imageData,
	}

	// Insert the recipe into the database
	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err = collection.InsertOne(ctx, recipe)
	if err != nil {
		http.Error(w, "Error saving the recipe", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func GetRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var recipe models.Recipe
	collection := database.GetCollection("recipes")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}
	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&recipe)

	if err != nil {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	// Encode the image to base64 to include in the response
	response := struct {
		ID          primitive.ObjectID `json:"id"`
		Name        string             `json:"name"`
		Category    string             `json:"category"`
		Ingredients []string           `json:"ingredients"`
		Steps       []string           `json:"steps"`
		Tags        []string           `json:"tags"`
		Summary     string             `json:"summary"`
		Image       string             `json:"image"` // Base64 encoded image
	}{
		ID:          recipe.ID,
		Name:        recipe.Name,
		Category:    recipe.Category,
		Ingredients: recipe.Ingredients,
		Steps:       recipe.Steps,
		Tags:        recipe.Tags,
		Summary:     recipe.Summary,
		Image:       base64.StdEncoding.EncodeToString(recipe.Image),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetRecipesByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["category"]

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"category": bson.M{"$regex": category, "$options": "i"}}

	// Start the find operation
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Could not retrieve recipes, check category", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var responseRecipes []map[string]interface{}

	for cursor.Next(ctx) {
		var recipe models.Recipe

		err := cursor.Decode(&recipe)
		if err != nil {
			http.Error(w, "Error decoding recipe", http.StatusInternalServerError)
			return
		}

		responseRecipe := map[string]interface{}{
			"id":          recipe.ID,
			"name":        recipe.Name,
			"category":    recipe.Category,
			"ingredients": recipe.Ingredients,
			"steps":       recipe.Steps,
			"tags":        recipe.Tags,
			"summary":     recipe.Summary,
			"image":       base64.StdEncoding.EncodeToString(recipe.Image),
		}

		responseRecipes = append(responseRecipes, responseRecipe)
	}

	if len(responseRecipes) == 0 {
		http.Error(w, "No recipes found for the given category", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responseRecipes)
}

func UpdateRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	var recipe models.Recipe
	_ = json.NewDecoder(r.Body).Decode(&recipe)

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	// Create a filter to match the document with the specific ObjectID
	filter := bson.M{"_id": objID}

	// Prepare the update data, wrapping it in $set to update only the fields provided
	update := bson.M{
		"$set": recipe,
	}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		http.Error(w, "Could not update recipe", http.StatusInternalServerError)
		return
	}

	if result.MatchedCount == 0 {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	// Retrieve the updated recipe to return
	err = collection.FindOne(ctx, filter).Decode(&recipe)
	if err != nil {
		http.Error(w, "Error while retrieving updated recipe", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func DeleteRecipe(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id := params["id"]

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		http.Error(w, "Invalid ID format", http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"_id": objID}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		http.Error(w, "Could not delete recipe", http.StatusInternalServerError)
		return
	}

	// Check if any document was actually deleted
	if result.DeletedCount == 0 {
		http.Error(w, "Recipe not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(
		map[string]string{
			"message": "Recipe was deleted successfully"})

}

func SearchRecipes(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	category := r.URL.Query().Get("category")

	if query == "" || category == "" {
		http.Error(w, "Both search query and category are required", http.StatusBadRequest)
		return
	}

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{
		"category": bson.M{"$regex": category, "$options": "i"},
		"$or": []bson.M{
			{"name": bson.M{"$regex": query, "$options": "i"}},
			{"ingredients": bson.M{"$regex": query, "$options": "i"}},
			{"tags": bson.M{"$regex": query, "$options": "i"}},
			{"summary": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Could no perform search", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var recipes []models.Recipe
	for cursor.Next(ctx) {
		var recipe models.Recipe

		err := cursor.Decode(&recipe)
		if err != nil {
			http.Error(w, "Error decoding search result", http.StatusInternalServerError)
			return
		}

		recipes = append(recipes, recipe)

	}

	if len(recipes) == 0 {
		http.Error(w, "No recipes found with your matching search and category", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
}
