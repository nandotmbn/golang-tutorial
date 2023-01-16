package controller_things

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
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var validate = validator.New()

var thingsCollection *mongo.Collection = configs.GetCollection(configs.DB, "things")

func RegisterThings() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var things models.Things
		defer cancel()
		c.BindJSON(&things)

		if validationErr := validate.Struct(&things); validationErr != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": validationErr.Error()})
			return
		}

		count, err_ := thingsCollection.CountDocuments(ctx, bson.M{"things_name": things.Thingname})

		if err_ != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": "Internal server error"})
			return
		}

		if count >= 1 {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Things name has been taken"})
			return
		}

		bytes, errors := bcrypt.GenerateFromPassword([]byte(things.Password), 14)
		if errors != nil {
			c.JSON(http.StatusBadRequest, bson.M{"data": "Password tidak valid"})
		}

		newThings := models.Things{
			Thingname: things.Thingname,
			Password:  string(bytes),
			CreatedAt: time.Now(),
		}

		result, err := thingsCollection.InsertOne(ctx, newThings)
		if err != nil {
			c.JSON(http.StatusInternalServerError, bson.M{"data": err.Error()})
			return
		}

		finalView := views.ThingsView{
			ThingsId:  result.InsertedID,
			Thingname: things.Thingname,
		}

		c.JSON(http.StatusCreated, bson.M{
			"status":  http.StatusCreated,
			"message": "success",
			"data":    finalView,
		})
	}
}
