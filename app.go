package main

import (
	"tutorial/configs"
	"tutorial/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	configs.ConnectDB()

	v1 := router.Group("/v1")

	routes.ThingsRoute(v1)
	routes.PointRoute(v1)

	router.Run("localhost:8080")
}
