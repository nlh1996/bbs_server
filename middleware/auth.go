package middleware

import (
	"bbs_server/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleWare 认证
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		headerToken := header[0]
		value,ok := common.TokenMap[headerToken]
		if (ok) {
			c.Request.Header["Authorization"][0] = value
			c.Next()
			return
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}

