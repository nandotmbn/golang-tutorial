package views

type MeterView struct {
	MeterId   interface{} `json:"meter_id,omitempty" validate:"required"`
	MeterName string      `json:"meter_name,omitempty" bson:"meter_name,omitempty" validate:"required,min=0"`
}

type PayloadRetriveId struct {
	MeterName string `json:"meter_name,omitempty" bson:"meter_name,omitempty" validate:"required,min=0"`
	Password  string `json:"password,omitempty" validate:"required,min=3,max=255"`
}

type FinalRetriveId struct {
	MeterId   interface{} `json:"meter_id,omitempty" bson:"_id,omitempty" validate:"required"`
	MeterName string      `json:"meter_name,omitempty" bson:"meter_name,omitempty" validate:"required,min=0"`
}
