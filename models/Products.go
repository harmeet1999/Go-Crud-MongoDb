package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Products struct {
	ID          primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name" bson:"name"`
	Description string             `json:"description" bson:"description"`
	Amount      string             `json:"amount" bson:"amount"`
}
