package views

import "time"

type PayloadPoint struct {
	Acidity     float32 `json:"acidity,omitempty"`
	Salinity    float32 `json:"salinity,omitempty"`
	Temperature float32 `json:"temperature,omitempty"`
}

type FinalPoint struct {
	Id          interface{} `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Acidity     float32     `json:"acidity,omitempty"`
	Salinity    float32     `json:"salinity,omitempty"`
	Temperature float32     `json:"temperature,omitempty"`
	CreatedAt   time.Time   `json:"created_at,omitempty" bson:"created_at,omitempty"`
}
