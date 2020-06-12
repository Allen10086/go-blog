package routers

import (
	"github.com/gin-gonic/gin"
	"blog/controller"
)


// 登录路由
func LoadLogin(e *gin.Engine) {
	LoginControl := &controller.LoginControl{}
	// 登录认证路由组
	LoginAuth := e.Group("/api/v1/user")
	{
		LoginAuth.POST("/login",LoginControl.Login)
	}
}



