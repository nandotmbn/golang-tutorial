package controller_point

import (
	"context"
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

var pointsCollection = configs.GetCollection(configs.DB, "points")
var thingsCollection = configs.GetCollection(configs.DB, "things")

func PointLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var pointPayload views.PayloadPoint
		defer cancel()
		var thingsId = c.Param("things_id")
		c.BindJSON(&pointPayload)

		if validationErr := validate.Struct(&pointPayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		thingsIdObj, err := primitive.ObjectIDFromHex(thingsId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		count, err_ := thingsCollection.CountDocuments(ctx, bson.M{"_id": thingsIdObj})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count == 0 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Things by given Id is not found"})
			return
		}

		newPoint := models.Point{
			ThingsId:  thingsIdObj,
			Latitude:  pointPayload.Latitude,
			Longitude: pointPayload.Longitude,
			Velocity:  pointPayload.Velocity,
		}

		result, err := pointsCollection.InsertOne(ctx, newPoint)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		finalView := views.FinalPoint{
			Id:        result.InsertedID,
			Latitude:  pointPayload.Latitude,
			Longitude: pointPayload.Longitude,
			Velocity:  pointPayload.Velocity,
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})

	}
}
