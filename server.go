package main

import (
	"server_go/api_warehouse_system/initializers"
	"server_go/api_warehouse_system/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

//var routes *gin.Engine

func main() {

	app := gin.Default()
	app.Use(cors.Default())
	routes.Setup(app)
	app.Run()

}
