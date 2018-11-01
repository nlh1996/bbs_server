package router

import (
	"bbs_server/controller/user"
	"bbs_server/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
	}
	
	v2 := router.Group("/v2")
	//v2群组使用中间件AuthMiddleWare
	v2.Use(middleware.AuthMiddleWare())
	{

	}
	router.Run(":8000")
}
