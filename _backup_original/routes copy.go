package routes

import (
	"github.com/gin-gonic/gin"
)

func Setup(app *gin.Engine) {

	//routes = gin.Default()
	app.GET("/api/*catch-all", func(c *gin.Context) {
		//path := strings.Split(c.Request.RequestURI, "/")

		//fmt.Println(c.Request.RequestURI)
		//fmt.Println(path[1])
		// c.IndentedJSON(http.StatusOK, gin.H{
		// 	"message": "api sub folder routes",
		// })
		Coredata(app)
		Apiquote(app)
	})

}
