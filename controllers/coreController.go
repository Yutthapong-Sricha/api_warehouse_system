package controllers

import (
	"fmt"
	"net/http"
	modelsCore "server_go/api_warehouse_system/models/apiCoreData"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Hello(c *gin.Context) {
	sID := uuid.NewString()

	//session := sessions.Default(c)

	session := sessions.Default(c)
	session.Set("id", sID)

	session.Set("email", "test@gmail.com")
	session.Save()
	fmt.Println(session.ID())
	fmt.Println(session.Get("id"))
	fmt.Println(session.Get("email"))
	// c.JSON(http.StatusOK, gin.H{
	// 	"message":   "Google UUID",
	// 	"sessionID": sID,
	// })

	// check session id
	/*if sessionID == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "unauthorized",
		})
		c.Abort()
	}*/
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
	c.IndentedJSON(http.StatusOK, gin.H{"data": branchs})

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
