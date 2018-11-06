package middleware

import (
	"fmt"
	"bbs_server/common"
	"net/http"

	"github.com/gin-gonic/gin"
)

// AuthMiddleWare cookie认证
func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header["Authorization"]
		headerToken := header[0]

		user, ok := common.TokenMap[headerToken]
		if (ok) {
			fmt.Println(user)
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

