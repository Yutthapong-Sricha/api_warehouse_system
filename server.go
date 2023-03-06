package main

import (
	"net/http"
	"server_go/api_warehouse_system/initializers"

	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
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

	router.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusForbidden, gin.H{"message": "error api"})
	})

	// start test session
	store := cookie.NewStore([]byte("spc"))
	router.Use(sessions.Sessions("session_id", store))

	router.GET("/hello", func(c *gin.Context) {
		session := sessions.Default(c)
		if session.Get("hello") != "world" {
			session.Set("hello", "world")
			session.Save()
		}

		c.JSON(http.StatusOK, gin.H{"hello": session.Get("hello")})
	})

	router.GET("/incr", func(c *gin.Context) {
		session := sessions.Default(c)
		var count int
		v := session.Get("count")
		if v == nil {
			count = 0
		} else {
			count = v.(int)
			count++
		}

		session.Set("count", count)
		session.Save()
		c.JSON(http.StatusOK, gin.H{"count": count})
	})

	// end test session

	//initializeRoutes() ใช้งานอยู่ comment เพราะ test session

	router.Run()

}
