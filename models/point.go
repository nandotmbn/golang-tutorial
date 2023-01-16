package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Point struct {
	ThingsId  primitive.ObjectID `json:"things_id,omitempty" bson:"things_id,omitempty" validate:"required,min=0"`
	Latitude  string             `json:"latitude,omitempty" validate:"required"`
	Longitude string             `json:"longitude,omitempty" validate:"required"`
	Velocity  int64              `json:"velocity,omitempty" validate:"required"`
	CreatedAt time.Time          `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
