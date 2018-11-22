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
		//获取到请求头中的token,在map中查找是否存在该token
		value,ok := common.TokenMap[headerToken]
		if (ok) {
			//将请求头中的token换成内存中相应的用户id,在之后的路由中不需要前端传用户后端可自行从请求头中获取用户。
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

