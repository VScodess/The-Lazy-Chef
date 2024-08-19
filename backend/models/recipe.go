package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Recipe struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name"`
	Category    string             `json:"category"`
	Ingredients []string           `json:"ingredients"`
	Steps       []string           `json:"steps"`
	Tags        []string           `json:"tags"`
	Summary     string             `json:"summary"`
	Image       []byte             `json:"image,omitempty" bson:"image,omitempty"`
}
