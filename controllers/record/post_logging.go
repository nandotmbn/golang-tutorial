package controller_point

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
	"tutorial/configs"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

var pointsCollection = configs.GetCollection(configs.DB, "record")
var meterCollection = configs.GetCollection(configs.DB, "meter")

func RecordLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var pointPayload views.PayloadPoint
		defer cancel()
		var meterId = c.Param("meter_id")
		c.BindJSON(&pointPayload)

		fmt.Println(pointPayload)

		if validationErr := validate.Struct(&pointPayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		meterIdObj, err := primitive.ObjectIDFromHex(meterId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		count, err_ := meterCollection.CountDocuments(ctx, bson.M{"_id": meterIdObj})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count == 0 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Things by given Id is not found"})
			return
		}

		newPoint := models.Record{
			MeterId:     meterIdObj,
			Acidity:     pointPayload.Acidity,
			Salinity:    pointPayload.Salinity,
			Temperature: pointPayload.Temperature,
			CreatedAt:   time.Now().UTC(),
		}

		result, err := pointsCollection.InsertOne(ctx, newPoint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		finalView := views.FinalPoint{
			Id:          result.InsertedID,
			Acidity:     pointPayload.Acidity,
			Salinity:    pointPayload.Salinity,
			Temperature: pointPayload.Temperature,
			CreatedAt:   time.Now(),
		}

		json_data, err__ := json.Marshal(finalView)
		if err__ != nil {
			log.Fatal("FUCJK")
		}

		http.Post("https://gdsc-pens-iot-listener-lxz6xwlfka-et.a.run.app/good-ponds/"+meterId, "application/json", bytes.NewBuffer(json_data))

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})

	}
}
