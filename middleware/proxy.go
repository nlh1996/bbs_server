package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// CrossDomain 跨域处理
func CrossDomain() gin.HandlerFunc {
	return func(c *gin.Context) {
		//origin := "http://"+c.Request.Host
		// 设置允许访问的域
		//c.Header("Access-Control-Allow-Origin", "http://115.159.77.155:11500")
		c.Header("Access-Control-Allow-Origin", "http://192.168.1.11:3001")
		// 服务器支持的所有跨域请求的方法,为了避免浏览次请求的多次'预检'请求
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE,UPDATE")
		// header的类型
		c.Header("Access-Control-Allow-Headers", "Authorization, Content-Length, X-CSRF-Token,Token,session,X_Requested_With,Accept, Origin, Host, Connection, Accept-Encoding, Accept-Language,DNT, X-CustomHeader, Keep-Alive, User-Agent, X-Requested-With, If-Modified-Since, Cache-Control, Content-Type, Pragma")
		// 跨域关键设置 让浏览器可以解析
		c.Header("Access-Control-Expose-Headers", "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers,Cache-Control,Content-Language,Content-Type,Expires,Last-Modified,Pragma,FooBar")
		// 本次预检请求的有效期，设置为24小时
		c.Header("Access-Control-Max-Age", "86400")
		// 跨域请求是否需要带cookie信息 设置为true
		c.Header("Access-Control-Allow-Credentials", "true")
		// 设置返回格式是json
		c.Set("content-type", "application/json")
		// 预请求回应
		if c.Request.Method == "OPTIONS" {
			c.String(http.StatusOK, "")
		}
		c.Next()
		return
	}
}
