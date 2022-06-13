package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Amazingzone struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Item   	string             `json:"item,omitempty"`
	Size 	string              `json:"size,omitempty"`
	Colour 	string				`json:"colour,omitempty"`
}
