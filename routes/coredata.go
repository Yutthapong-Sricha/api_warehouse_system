package routes

import (
	"server_go/api_warehouse_system/controllers"

	"github.com/gin-gonic/gin"
)

func Coredata(routes *gin.Engine) {
	coredata := routes.Group("/coredata")
	{
		coredata.GET("/", controllers.Hello)                      // localhost:xxxx/coredata/
		coredata.GET("/position", controllers.Positions)          // localhost:xxxx/coredata/position
		coredata.GET("/getposition/:id", controllers.GetPosition) // localhost:xxxx/coredata/getposition/xx

		coredata.GET("/branchs", controllers.Branchs)         // localhost:xxxx/coredata/branchs
		coredata.GET("/getbranch/:id", controllers.GetBranch) // localhost:xxxx/coredata/getbranch/xx
	}
}
