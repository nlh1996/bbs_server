package router

import (
	"bbs_server/controller/index"
	"bbs_server/controller/thread"
	"bbs_server/controller/user"
	"bbs_server/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() {
	router := gin.Default()
	router.Use(middleware.CrossDomain())
	v1 := router.Group("/v1")
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
		v1.GET("/index", index.GetInfo)
	}

	v2 := router.Group("/v2")
	//v2群组使用中间件AuthMiddleWare
	v2.Use(middleware.AuthMiddleWare())
	{
		v2.POST("/publish", thread.Publish)
		v2.POST("/isload", user.IsLoad)
	}
	router.Run(":8000")
}
