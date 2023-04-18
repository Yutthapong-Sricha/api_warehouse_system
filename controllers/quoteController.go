package controllers

import (
	"net/http"
	modelsQuote "server_go/api_warehouse_system/models/apiQuote"

	"github.com/gin-gonic/gin"
)

func Hello_quote(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, gin.H{
		"message": "has api Group hello quote",
	})
}

func Promotions(c *gin.Context) {
	act := c.Param("act")
	promotions := modelsQuote.ListPromotion(act)
	c.IndentedJSON(http.StatusOK, promotions)
}

func GetPromotion(c *gin.Context) {
	id_enc := c.Param("id_enc")
	promotion := modelsQuote.GetPromotion(id_enc)
	c.IndentedJSON(http.StatusOK, promotion)
}

func AddShortCust(c *gin.Context) {
	ret, err := modelsQuote.InstShortCust(c)
	//fmt.Println(err.Error())
	if err != nil {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "error", "message": err.Error()})
	} else {
		c.IndentedJSON(http.StatusOK, gin.H{"status": "success", "message": "บันทึก " + ret + " เรียบร้อยแล้ว"})
	}

}
