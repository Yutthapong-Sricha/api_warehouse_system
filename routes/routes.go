package routes

import (
	"fmt"
	"net/http"
	"strings"

	"server_go/api_warehouse_system/controllers"
	"server_go/api_warehouse_system/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {

	//app.GET("/api/*catch-all", func(c *gin.Context) {
	//path := strings.Split(c.Request.RequestURI, "/")

	//fmt.Println(c.Request.RequestURI)
	//fmt.Println(path[1])
	// c.IndentedJSON(http.StatusOK, gin.H{
	// 	"message": "api sub folder routes",
	// })

	//})

	app.GET("/landing", controllers.Landing)
	app.POST("/login", controllers.Login)
	app.POST("/logout", controllers.Logout)

	app.GET("/header", func(c *gin.Context) {
		fmt.Println(c.Request.UserAgent())
		Authorization := c.Request.Header.Get("Authorization")
		bearerToken := strings.TrimPrefix(Authorization, "Bearer ")
		fmt.Println(bearerToken)
		mysession := sessions.Default(c)
		sess_value := mysession.Get(bearerToken)
		fmt.Println(sess_value)
	})

	//

	app.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "path api ",
		})
	})
	api := app.Group("/api")
	//api.Use(cors.Default())
	api.Use(middleware.AuthToken())
	{

		// api.GET("/header", func(c *gin.Context) {
		// 	fmt.Println(c.Request.UserAgent())
		// })
		Coredata(api)
		Apiquote(api)
		Search(api)
	}

}
