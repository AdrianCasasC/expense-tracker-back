package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type ExpenseDto struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Value    float64            `json:"value" bson:"value"`
	Type     string             `json:"type" bson:"type"`
	Category string             `json:"category" bson:"category"`
	Date     string             `json:"date" bson:"date"`
}

type ExpenseEntity struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name"`
	Value    float64            `json:"value" bson:"value"`
	Category string             `json:"category" bson:"category"`
	Date     string             `json:"date" bson:"date"`
}
