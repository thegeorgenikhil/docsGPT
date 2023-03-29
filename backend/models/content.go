package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Content struct {
	ID         primitive.ObjectID `json:"_id" bson:"_id"`
	DocumentId primitive.ObjectID `json:"document_id" bson:"document_id"`
	Content    string             `json:"content" bson:"content"`
	CreatedAt  time.Time          `json:"createdAt" bson:"created_at"`
	UpdatedAt  time.Time          `json:"updatedAt" bson:"updated_at"`
}
