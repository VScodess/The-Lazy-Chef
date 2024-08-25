package handlers

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"The-Lazy-Chef/backend/config"
	"The-Lazy-Chef/backend/database"
	"The-Lazy-Chef/backend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// fetchRecipes fetches all recipes from the database.
func fetchRecipes(cfg *config.Config) ([]map[string]interface{}, error) {
	var recipes []map[string]interface{}

	collection, ctx, cancel := getCollectionWithContext(cfg)
	defer cancel()

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
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

	return recipes, cursor.Err()
}

// parseRecipeForm parses the form data for creating or updating a recipe.
func parseRecipeForm(r *http.Request) (models.Recipe, error) {
	err := r.ParseMultipartForm(10 << 20)
	if err != nil {
		return models.Recipe{}, err
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		return models.Recipe{}, err
	}
	defer file.Close()

	imageData, err := ioutil.ReadAll(file)
	if err != nil {
		return models.Recipe{}, err
	}

	recipe := models.Recipe{
		ID:          primitive.NewObjectID(),
		Name:        r.FormValue("name"),
		Category:    r.FormValue("category"),
		Ingredients: cleanAndSplit(r.FormValue("ingredients")),
		Steps:       cleanAndSplit(r.FormValue("steps")),
		Tags:        cleanAndSplit(r.FormValue("tags")),
		Summary:     r.FormValue("summary"),
		Image:       imageData,
	}

	return recipe, nil
}

// insertRecipe inserts a new recipe into the database.
func insertRecipe(cfg *config.Config, recipe models.Recipe) error {
	collection, ctx, cancel := getCollectionWithContext(cfg)
	defer cancel()

	_, err := collection.InsertOne(ctx, recipe)
	return err
}

// fetchRecipeByID fetches a single recipe by its ID.
func fetchRecipeByID(cfg *config.Config, id string) (models.Recipe, error) {
	var recipe models.Recipe

	collection, ctx, cancel := getCollectionWithContext(cfg)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return models.Recipe{}, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&recipe)
	if err != nil {
		return models.Recipe{}, err
	}

	// Convert the image to base64 for the response if necessary
	recipe.Image = []byte(base64.StdEncoding.EncodeToString(recipe.Image))

	return recipe, nil
}

// updateRecipeByID updates an existing recipe in the database by its ID.
func updateRecipeByID(cfg *config.Config, id string, updatedRecipe models.Recipe) error {
	collection, ctx, cancel := getCollectionWithContext(cfg)
	defer cancel()

	objID, _ := primitive.ObjectIDFromHex(id)

	filter := bson.M{"_id": objID}
	update := bson.M{"$set": updatedRecipe}

	_, err := collection.UpdateOne(ctx, filter, update)
	return err
}

// deleteRecipeByID deletes a recipe from the database by its ID.
func deleteRecipeByID(cfg *config.Config, id string) error {
	collection, ctx, cancel := getCollectionWithContext(cfg)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

// fetchRecipesByCategory fetches recipes by their category.
func fetchRecipesByCategory(cfg *config.Config, category string) ([]models.Recipe, error) {
	var recipes []models.Recipe

	collection, ctx, cancel := getCollectionWithContext(cfg)
	defer cancel()

	filter := bson.M{"category": bson.M{"$regex": category, "$options": "i"}}
	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var recipe models.Recipe
		if err := cursor.Decode(&recipe); err != nil {
			return nil, err
		}
		recipes = append(recipes, recipe)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return recipes, nil
}

// searchRecipesByCategoryAndQuery performs a search within a category based on a query.
func searchRecipesByCategoryAndQuery(cfg *config.Config, category, query string) ([]models.Recipe, error) {
	var recipes []models.Recipe

	collection, ctx, cancel := getCollectionWithContext(cfg)
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
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var recipe models.Recipe

		err := cursor.Decode(&recipe)
		if err != nil {
			return nil, err
		}

		recipes = append(recipes, recipe)
	}

	return recipes, nil
}

// cleanAndSplit is a helper function that trims and splits a comma-separated string into a slice.
func cleanAndSplit(input string) []string {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return nil
	}
	return strings.Split(trimmed, ",")
}

// respondWithJSON is a helper function to send a JSON response.
func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func getCollectionWithContext(cfg *config.Config) (*mongo.Collection, context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	collection := database.GetCollection(cfg.CollectionName)
	return collection, ctx, cancel
}
