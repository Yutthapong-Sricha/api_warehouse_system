package routes

import (
	ctrlSearch "server_go/api_warehouse_system/controllers/search"

	"github.com/gin-gonic/gin"
)

func Search(route *gin.RouterGroup) {
	search := route.Group("/search")
	{
		search.POST("/cust", ctrlSearch.Searchcust)
		search.POST("/prod", ctrlSearch.Searchprod)
	}
}
