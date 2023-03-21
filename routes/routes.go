package routes

import (
	"net/http"

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
	api := app.Group("/api")
	app.GET("/", func(c *gin.Context) {
		c.IndentedJSON(http.StatusOK, gin.H{
			"message": "path api ",
		})
	})
	Coredata(api)
	Apiquote(app)
}
