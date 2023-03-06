package controllers

import (
	"net/http"
	models "server_go/api_warehouse_system/models/apiCoreData"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {

	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "has api Group",
	})
}

func Positions(c *gin.Context) {
	positions := models.ListPosition()
	c.IndentedJSON(http.StatusOK, positions)
}

func GetPosition(c *gin.Context) {
	id := c.Param("id")
	position := models.GetPosition(id)
	c.IndentedJSON(http.StatusOK, position)
}
