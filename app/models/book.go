package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Book struct {
	ID          primitive.ObjectID `bson:"_id" json:"id,omitempty"`
	Image       string             `json:"image"`
	Title       string             `json:"title"`
	Description string             `json:"description"`
	Author      string             `json:"author"`
	Rating      float64            `json:"rating"`
	Pages       int                `json:"pages"`
	Languages   string             `json:"languages"`
	Date        int                `json:"date"`
	Category    string             `json:"category"`
	Stock       int                `json:"stock"`
	ValidUntil  time.Time          `json:"validUntil"`
}
