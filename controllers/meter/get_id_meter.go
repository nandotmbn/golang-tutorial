package controller_things

import (
	"context"
	"net/http"
	"time"
	"tutorial/models"
	"tutorial/views"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"golang.org/x/crypto/bcrypt"
)

func GetIdMeter() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var thingPayload views.PayloadRetriveId
		defer cancel()
		c.BindJSON(&thingPayload)

		if validationErr := validate.Struct(&thingPayload); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		var resultThings models.Meter
		var finalPayload views.FinalRetriveId
		result := thingsCollection.FindOne(ctx, bson.M{"meter_name": thingPayload.MeterName})
		result.Decode(&resultThings)
		result.Decode(&finalPayload)
		err := bcrypt.CompareHashAndPassword([]byte(resultThings.Password), []byte(thingPayload.Password))
		if err != nil {
			c.JSON(http.StatusBadRequest, bson.M{
				"status":  http.StatusBadRequest,
				"message": "Bad request",
				"data":    "Things Name or Password is not valid",
			})
			return
		}

		c.JSON(http.StatusOK,
			bson.M{
				"status":  http.StatusOK,
				"message": "Success",
				"data":    finalPayload,
			},
		)
	}
}
