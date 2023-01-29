package routes

import (
	controller_record "tutorial/controllers/record"

	"github.com/gin-gonic/gin"
)

func RecordRoute(router *gin.RouterGroup) {
	router.GET("/record/:meter_id", controller_record.AllLogging())
	router.GET("/record/:meter_id/last", controller_record.LastLogging())
	router.POST("/record/:meter_id", controller_record.RecordLogging())
}
