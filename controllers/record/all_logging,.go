package controller_point

import (
	"context"
	"fmt"
	"net/http"
	"time"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func AllLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		var meterId = c.Param("meter_id")

		meterIdObj, err := primitive.ObjectIDFromHex(meterId)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		fmt.Println(meterIdObj)

		results, err_ := pointsCollection.Find(ctx, bson.M{"meter_id": meterIdObj})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		defer results.Close(ctx)

		var roles []views.PayloadPoint
		for results.Next(ctx) {
			var singleRoles views.PayloadPoint
			if err = results.Decode(&singleRoles); err != nil {
				c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			}

			roles = append(roles, singleRoles)
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusOK,
			"message": "success",
			"data":    roles,
		})

	}
}
