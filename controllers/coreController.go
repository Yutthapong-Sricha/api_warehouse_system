package controllers

import (
	"net/http"
	modelsCore "server_go/api_warehouse_system/models/apiCoreData"

	"github.com/gin-gonic/gin"
)

func Hello(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "has api Group",
	})
}

func Positions(c *gin.Context) {
	act := c.Param("act")
	positions := modelsCore.ListPosition(act)
	c.IndentedJSON(http.StatusOK, positions)
}

func GetPosition(c *gin.Context) {
	id_enc := c.Param("id_enc")
	position := modelsCore.GetPosition(id_enc)
	c.IndentedJSON(http.StatusOK, position)
}

func Branchs(c *gin.Context) {
	act := c.Param("act")
	branchs := modelsCore.ListBranch(act)
	c.IndentedJSON(http.StatusOK, branchs)
}

func GetBranch(c *gin.Context) {
	id_enc := c.Param("id_enc")
	branch := modelsCore.GetBranch(id_enc)
	c.IndentedJSON(http.StatusOK, branch)
}

func Categorys(c *gin.Context) {
	act := c.Param("act")
	categorys := modelsCore.ListCategory(act)
	c.IndentedJSON(http.StatusOK, categorys)
}

func GetCategory(c *gin.Context) {
	id_enc := c.Param("id_enc")
	category := modelsCore.GetCategory(id_enc)
	c.IndentedJSON(http.StatusOK, category)
}

func Suppliers(c *gin.Context) {
	act := c.Param("act")
	suppliers := modelsCore.ListSupplier(act)
	c.IndentedJSON(http.StatusOK, suppliers)
}

func Getsupplier(c *gin.Context) {
	id_enc := c.Param("id_enc")
	supplier := modelsCore.GetSupplier(id_enc)
	c.IndentedJSON(http.StatusOK, supplier)
}

func Products(c *gin.Context) {
	act := c.Param("act")
	products := modelsCore.ListProduct(act)
	c.IndentedJSON(http.StatusOK, products)
}

func Getproduct(c *gin.Context) {
	id_enc := c.Param("id_enc")
	product := modelsCore.GetProduct(id_enc)
	c.IndentedJSON(http.StatusOK, product)
}
