package middleware

import (
	"net/http"
	WSession "server_go/api_warehouse_system/models/sessions"

	"github.com/gin-gonic/gin"
)

// func AuthToken(w http.ResponseWriter, r *http.Request) gin.HandlerFunc {
func AuthToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// c.Header("Access-Control-Allow-Origin", "http://localhost:3000")
		// c.Header("Access-Control-Allow-Credentials", "true")
		// c.Header("Access-Control-Allow-Headers", "*")

		// if c.Request.Method == "OPTIONS" {
		// 	c.AbortWithStatus(204)
		// }
		err := WSession.SessionAuth(c)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": err.Error()})
		}
		c.Next()
	}
}
