package config

import (
	"log"
	"os"
)

type Config struct {
	MongoURI     string
	DatabaseName string
	CollectionName string
	ServerPort   string
}

// LoadConfig loads configuration from environment variables or defaults
func LoadConfig() *Config {
	config := &Config{
		MongoURI:       getEnv("MONGO_URI", "mongodb://localhost:27017"),
		DatabaseName:   getEnv("DB_NAME", "lazy_chef"),
		CollectionName: getEnv("COLLECTION_NAME", "recipes"),
		ServerPort:     getEnv("SERVER_PORT", "8080"),
	}

	log.Printf("Config Loaded: %+v\n", config)
	return config
}

// getEnv gets an environment variable or returns a default value
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
