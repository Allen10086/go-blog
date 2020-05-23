package router

import (
	"blog/middleware"
	"github.com/gin-gonic/gin"
	"blog/controller"
)



// 登录路由
func LoginRouter()  {
	router := gin.Default()
	// 中间件解决跨域问题
	router.Use(middleware.Cors())

	LoginControl := &controller.LoginControl{}
	// 登录认证路由组
	LoginAuth := router.Group("/api/user")
	{
		LoginAuth.POST("/login",LoginControl.Iogin)
	}

	router.Run(":8080")

}
