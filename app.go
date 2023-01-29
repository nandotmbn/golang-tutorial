package main

import (
	"net/http"
	"tutorial/configs"
	"tutorial/routes"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	//check origin will check the cross region source (note : please not using in production)
	CheckOrigin: func(r *http.Request) bool {
		//Here we just allow the chrome extension client accessable (you should check this verify accourding your client source)
		return true
	},
}

func main() {
	router := gin.Default()
	configs.ConnectDB()

	v1 := router.Group("/v1")

	// routes_websocket.PointWebsocketRoute(v1)
	routes.MeterRoute(v1)
	routes.RecordRoute(v1)

	router.Run(":8080")
}
