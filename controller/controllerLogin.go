package controller

import (
	"blog/dao/user"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginControl struct {
}

type Login struct {
	UserName string `form:"username" json:"username"`
	Password string `form:"password" json:"password"`
}

// 登录认证control
/*
Go语言中的方法（Method）是一种作用于特定类型变量的函数。这种特定类型变量叫做接收者（Receiver）。接收者的概念就类似于其他语言中的this或者 self。
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
LoginControl  实现了Login方法
*/
func (w *LoginControl) Login(c *gin.Context) {
	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	var login Login
	// 如果有值没有错误返回json
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 调用查询函数查询数据库中的用户名和密码
	_, err = user.GetUserInfo(login.UserName, login.Password)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 登录失败返回code 1001
			"code":    1001,
			"message": "failed",
		})
		return
	}

	// 没有数据库的用下面这个方法：这里先写死账号和密码  有数据库的要从数据库中获取
	//if (login.UserName != "code" || login.Password != "123456") {
	//	c.JSON(http.StatusOK, gin.H{
	//		// 登录失败返回code 1001
	//		"code": 1001,
	//		"message":"failed",
	//	})
	//	return
	//}

	c.JSON(http.StatusOK, gin.H{
		// 登录失败返回code 1000
		"code":    1000,
		"message": "success",
	})
	return
}
