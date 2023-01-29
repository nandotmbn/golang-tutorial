package controller_point

import (
	"context"
	"net/http"
	"time"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func LastLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var meterId = c.Param("meter_id")

		meterIdObj, err := primitive.ObjectIDFromHex(meterId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		var roles views.PayloadPoint
		myOptions := options.FindOne()
		myOptions.SetSort(bson.M{"$natural": -1})
		results := pointsCollection.FindOne(ctx, bson.M{"meter_id": meterIdObj}, myOptions)
		results.Decode(&roles)

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusOK,
			"message": "success",
			"data":    roles,
		})

	}
}
