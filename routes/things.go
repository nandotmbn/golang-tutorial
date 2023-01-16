package routes

import (
	controller_things "tutorial/controllers/things"

	"github.com/gin-gonic/gin"
)

func ThingsRoute(router *gin.RouterGroup) {
	router.POST("/things", controller_things.RegisterThings())
	router.POST("/things/retriveid", controller_things.GetIdThings())
}
