package search

import (
	"net/http"
	modelsSearch "server_go/api_warehouse_system/models/apiSearch"

	"github.com/gin-gonic/gin"
)

func Searchcust(c *gin.Context) {
	listSearch := modelsSearch.Cust(c)
	c.IndentedJSON(http.StatusOK, listSearch)
	// c.IndentedJSON(200, gin.H{
	// 	"status":  "success",
	// 	"message": "Success Search",
	// 	"data":    listSearch,
	// })
}

func Searchprod(c *gin.Context) {
	listSearch := modelsSearch.Prod(c)
	c.IndentedJSON(http.StatusOK, listSearch)
}
