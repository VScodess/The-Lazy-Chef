package handlers

import (
	"The-Lazy-Chef/backend/database"
	"The-Lazy-Chef/backend/models"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Recipe struct with struct tags, provides metadata about struct fields
// helps with marshalling and umarhsalling data. Think ID vs id

func GetRecipes(w http.ResponseWriter, r *http.Request) {
	var recipes []models.Recipe
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
		recipes = append(recipes, recipe)
	}

	if err := cursor.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)

}

func CreateRecipe(w http.ResponseWriter, r *http.Request) {
	var recipe models.Recipe
	_ = json.NewDecoder(r.Body).Decode(&recipe)

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	result, err := collection.InsertOne(ctx, recipe)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Retrieve the inserted ID and add it to the response
	insertedID := result.InsertedID.(primitive.ObjectID)

	// Return the newly created recipe along with its MongoDB-generated _id
	response := struct {
		ID primitive.ObjectID `json:"id"`
		models.Recipe
	}{
		ID:     insertedID,
		Recipe: recipe,
	}

	// Encode the response as JSON and send it back to the client
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipe)
}

func GetRecipesByCategory(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	category := params["category"]

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"category": category}

	// Start the find operation
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		http.Error(w, "Could not retrieve recipes, check category", http.StatusInternalServerError)
		return
	}
	defer cursor.Close(ctx)

	var recipes []models.Recipe

	for cursor.Next(ctx) {
		var recipe models.Recipe

		err := cursor.Decode(&recipe)
		if err != nil {
			http.Error(w, "Error decoding recipe", http.StatusInternalServerError)
			return
		}

		recipes = append(recipes, recipe)
	}

	if len(recipes) == 0 {
		http.Error(w, "No recipes found for the given category", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(recipes)
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
