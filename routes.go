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

	// Simple group: coredata
	coredata := router.Group("/coredata")
	{
		coredata.GET("/", controllers.Hello)             // localhost:xxxx/coredata/
		coredata.GET("/position", controllers.Positions) // localhost:xxxx/coredata/position

		coredata.GET("/getposition/:id", controllers.GetPosition) // localhost:xxxx/coredata/getposition/xx
	}

}
