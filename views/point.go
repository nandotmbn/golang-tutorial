package views

type PayloadPoint struct {
	Latitude  string `json:"latitude,omitempty" validate:"required"`
	Longitude string `json:"longitude,omitempty" validate:"required"`
	Velocity  int64  `json:"velocity,omitempty" validate:"required"`
}

type FinalPoint struct {
	Id        interface{} `json:"_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Latitude  string      `json:"latitude,omitempty" validate:"required"`
	Longitude string      `json:"longitude,omitempty" validate:"required"`
	Velocity  int64       `json:"velocity,omitempty" validate:"required"`
}
