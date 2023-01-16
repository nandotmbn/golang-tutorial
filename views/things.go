package views

type ThingsView struct {
	ThingsId  interface{} `json:"things_id,omitempty" validate:"required"`
	Thingname string      `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
}

type PayloadRetriveId struct {
	Thingname string `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
	Password  string `json:"password,omitempty" validate:"required,min=3,max=255"`
}

type FinalRetriveId struct {
	ThingsId  interface{} `json:"things_id,omitempty" bson:"_id,omitempty" validate:"required"`
	Thingname string      `json:"things_name,omitempty" bson:"things_name,omitempty" validate:"required,min=0"`
}
