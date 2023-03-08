package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Apiquote(routes *gin.Engine) {
	quotation := routes.Group("/quotation")
	{
		quotation.GET("/", func(c *gin.Context) {
			c.IndentedJSON(http.StatusOK, gin.H{
				"message": "api quotation",
			})
		}) // localhost:xxxx/apistaff/
	}

}
