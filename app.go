package main

import (
	// "tutorial/function"
	// multiple "tutorial/function/multiple"
	// "tutorial/function/substract"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.JSON(200, "Hello World")
	})

	router.Run("localhost:6000")
}
