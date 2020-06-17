package main

import (
	"blog/database"
	"blog/middleware"
	"blog/models"
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
	// 模型和数据库表映射 创建表
	database.DB.AutoMigrate(&models.UserInfo{},&models.ArticleCategory{},&models.ArticleList{})
	// 注册路由
	r := gin.Default()
	// 将中间件注册到全局 对所有路由生效
	r.Use(middleware.Cors())
	// 注册登录路由
	routers.LoadLogin(r)
	// 文章分类路由
	routers.LoadCategory(r)
	// 图片上传
	routers.ImgUpload(r)
	// 文章管理
	routers.LoadArticle(r)

	r.Run(":8081")

}
