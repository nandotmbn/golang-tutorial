package models

import (
	"time"
)

type Meter struct {
	MeterName string    `json:"meter_name,omitempty" bson:"meter_name,omitempty" validate:"required,min=0"`
	Password  string    `json:"password,omitempty" validate:"required,min=3,max=255"`
	CreatedAt time.Time `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
