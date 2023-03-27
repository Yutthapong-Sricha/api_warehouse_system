package routes

import (
	"server_go/api_warehouse_system/controllers"

	"github.com/gin-gonic/gin"
)

func Apiquote(routes *gin.RouterGroup) {
	quotation := routes.Group("/quotation")
	{
		quotation.GET("/", controllers.Hello_quote) // localhost:xxxx/apistaff/

		quotation.GET("/promotions/*act", controllers.Promotions)
		quotation.GET("/getpromotion/:id_enc", controllers.GetPromotion)
	}

}
