package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Record struct {
	MeterId     primitive.ObjectID `json:"meter_id,omitempty" bson:"meter_id,omitempty" validate:"required,min=0"`
	Acidity     float32            `json:"acidity,omitempty"`
	Salinity    float32            `json:"salinity,omitempty"`
	Temperature float32            `json:"temperature,omitempty"`
	CreatedAt   time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
