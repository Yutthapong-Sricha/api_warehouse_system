package controllers

import (
	"net/http"
	models "server_go/api_warehouse_system/models/apiCoreData"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "Group v1",
	})
}

func Positions(c *gin.Context) {
	Positions := models.ListPosition()
	c.IndentedJSON(http.StatusOK, Positions)
}

func GetPosition(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "GetPosition",
	})
}
