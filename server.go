package main

import (
	"server_go/api_warehouse_system/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

var router *gin.Engine

func main() {

	router = gin.Default()
	router.Use(cors.Default())

	initializeRoutes()

	router.Run()

}
