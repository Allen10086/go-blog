package main

import (
	"blog/database"
	"blog/middleware"
	"blog/routers"
	"github.com/gin-gonic/gin"
)

func main() {
	// 数据库初始化
	err := database.InitMySql()
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer database.DB.Close()

	// 注册路由
	r := gin.Default()
	// 将中间件注册到全局 对所有路由生效
	r.Use(middleware.Cors())
	// 注册登录路由
	routers.LoadLogin(r)
	// 文章分类路由
	routers.LoadCategory(r)
	r.Run(":8081")

}
