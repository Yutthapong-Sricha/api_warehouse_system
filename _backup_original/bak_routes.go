package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func initializeRoutes() {

	// router.GET("/", func(c *gin.Context) {
	// 	c.IndentedJSON(http.StatusForbidden, gin.H{"message": "error api"})
	// })

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "api main",
		})
	})

	// Simple group: coredata
	// coredata := router.Group("/coredata")
	// {
	// 	coredata.GET("/", controllers.Hello)                      // localhost:xxxx/coredata/
	// 	coredata.GET("/position", controllers.Positions)          // localhost:xxxx/coredata/position
	// 	coredata.GET("/getposition/:id", controllers.GetPosition) // localhost:xxxx/coredata/getposition/xx

	// 	coredata.GET("/branchs", controllers.Branchs)         // localhost:xxxx/coredata/branchs
	// 	coredata.GET("/getbranch/:id", controllers.GetBranch) // localhost:xxxx/coredata/getbranch/xx
	// }

}
