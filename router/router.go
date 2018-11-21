package router

import (
	"bbs_server/controller/index"
	"bbs_server/controller/post"
	"bbs_server/controller/user"
	"bbs_server/middleware"

	"github.com/gin-gonic/gin"
)

// Init 初始化路由
func Init() {
	router := gin.Default()
	//CrossDomain跨域处理，options请求处理
	router.Use(middleware.CrossDomain())
	// v1群组对任何人开放
	v1 := router.Group("/v1")
	{
		v1.POST("/login", user.Login)
		v1.POST("/register", user.Register)
		v1.GET("/index", index.GetInfo)
		v1.GET("/posts", post.GetPosts)
		v1.GET("/post", post.GetPost)
	}

	v2 := router.Group("/v2")
	//v2群组使用中间件AuthMiddleWare，需要token权限才能请求到
	v2.Use(middleware.AuthMiddleWare())
	{
		v2.POST("/publish", post.Publish)
		v2.POST("/isload", user.IsLoad)
		v2.POST("/reply1", post.Reply1)
	}
	router.Run(":8000")
}
