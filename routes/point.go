package routes

import (
	controller_point "tutorial/controllers/point"

	"github.com/gin-gonic/gin"
)

func PointRoute(router *gin.RouterGroup) {
	router.POST("/point/:things_id", controller_point.PointLogging())
}
