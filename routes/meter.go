package routes

import (
	controller_meter "tutorial/controllers/meter"

	"github.com/gin-gonic/gin"
)

func MeterRoute(router *gin.RouterGroup) { // http://localhost:8080/v1
	router.POST("/meter", controller_meter.RegisterMeter()) // http://localhost:8080/v1/meter
	router.POST("/meter/retriveid", controller_meter.GetIdMeter())
}
