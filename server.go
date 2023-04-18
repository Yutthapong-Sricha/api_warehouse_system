package main

import (
	"os"
	"server_go/api_warehouse_system/initializers"
	"server_go/api_warehouse_system/routes"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
}

//var routes *gin.Engine

func main() {

	app := gin.Default()
	secretKey := os.Getenv("SECRETKEY")
	store := cookie.NewStore([]byte(secretKey))
	app.Use(sessions.Sessions("sessionW", store))

	//app.Use(cors.Default())
	app.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		//ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
		// AllowOriginFunc: func(origin string) bool {
		// 	return origin == "http://localhost:3000"
		// },
		MaxAge: 12 * time.Hour,
	}))

	routes.Setup(app)
	app.Run()

}
