package main

import (
	"The-Lazy-Chef/backend/database"
	"The-Lazy-Chef/backend/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func cleanAndSplit(input string) []string {
	parts := strings.Split(input, ",")
	for i, part := range parts {
		parts[i] = strings.Trim(part, "' ")
	}
	return parts
}

func main() {

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error getting current working directory:", err)
		return
	}
	fmt.Println("Current working directory:", wd)

	database.Connect()

	jsonFile, err := os.Open("../data/recipes.json")
	if err != nil {
		log.Fatal("Error opening JSON file: ", err)
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var recipes []struct {
		Name        string `json:"name"`
		Tags        string `json:"tags"`
		Steps       string `json:"steps"`
		Description string `json:"description"`
		Ingredients string `json:"ingredients"`
		Category    string `json:"category"`
		ImageBytes  string `json:"image_bytes"`
	}

	err = json.Unmarshal(byteValue, &recipes)
	if err != nil {
		log.Fatal("Error parsing JSON: ", err)
	}

	collection := database.GetCollection("recipes")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for _, recipe := range recipes {

		imageData, err := base64.StdEncoding.DecodeString(recipe.ImageBytes)
		if err != nil {
			log.Fatal("Error decoding image data:", err)
		}

		ingredients := cleanAndSplit(recipe.Ingredients)
		steps := cleanAndSplit(recipe.Steps)
		tags := cleanAndSplit(recipe.Tags)

		recipeModel := models.Recipe{
			ID:          primitive.NewObjectID(),
			Name:        recipe.Name,
			Category:    recipe.Category,
			Ingredients: ingredients,
			Steps:       steps,
			Tags:        tags,
			Summary:     recipe.Description,
			Image:       imageData,
		}

		_, err = collection.InsertOne(ctx, recipeModel)

		if err != nil {
			log.Fatal("Error inserting the recipe into MongoDB:", err)
		}
	}
}
