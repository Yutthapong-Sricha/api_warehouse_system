package main

import (
	"net/http"

	"server_go/api_warehouse_system/controllers"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "api main",
		})
	})

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/", controllers.Hello)
		v1.GET("/position", controllers.Positions)

		v1.GET("/getposition/:id", controllers.GetPosition)
		//v1.POST("/login", loginEndpoint)
		//v1.POST("/submit", submitEndpoint)
		//v1.POST("/read", readEndpoint)
	}

	// Simple group: v2
	//v2 := router.Group("/v2")
	//{
	//v2.POST("/login", loginEndpoint)
	//v2.POST("/submit", submitEndpoint)
	//v2.POST("/read", readEndpoint)
	//}
}
