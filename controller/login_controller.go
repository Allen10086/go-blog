package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginControl struct {

}

// 登录认证control
func (w *LoginControl) Iogin(c *gin.Context) {
	// ShouldBind()会根据请求的Content-Type自行选择绑定器
	var login Login
	// 如果有值没有错误返回json
	err := c.ShouldBind(&login)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// 这里先写死账号和密码  后期要从数据库中获取
	if (login.User != "code" || login.Passwd != "123456") {
		c.JSON(http.StatusOK, gin.H{
			// 登录失败返回code 2000
			"code": 2000,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		// 登录失败返回code 1000
		"code": 1000,
	})
	return
}