package routers

import (
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"blog/controller"
)


// 登录路由
func SetupRouter() *gin.Engine {
	router := gin.Default()
	// 将中间件注册到全局 对所有路由生效
	router.Use(middleware.Cors())

	LoginControl := &controller.LoginControl{}
	// 登录认证路由组
	LoginAuth := router.Group("/api/user")
	{
		LoginAuth.POST("/login",LoginControl.Iogin)
	}
	return router
}
