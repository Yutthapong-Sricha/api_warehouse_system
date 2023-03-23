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

func Branchs(c *gin.Context) {
	branchs := models.ListBranch()
	c.IndentedJSON(http.StatusOK, branchs)
}

func GetBranch(c *gin.Context) {
	id := c.Param("id")
	branch := models.GetBranch(id)
	c.IndentedJSON(http.StatusOK, branch)
}

func Categorys(c *gin.Context) {
	categorys := models.ListCategory()
	c.IndentedJSON(http.StatusOK, categorys)
}

func GetCategory(c *gin.Context) {
	id := c.Param("id")
	category := models.GetCategory(id)
	c.IndentedJSON(http.StatusOK, category)
}

func Suppliers(c *gin.Context) {
	suppliers := models.ListSupplier()
	c.IndentedJSON(http.StatusOK, suppliers)
}

func Getsupplier(c *gin.Context) {
	id := c.Param("id")
	supplier := models.GetSupplier(id)
	c.IndentedJSON(http.StatusOK, supplier)
}
