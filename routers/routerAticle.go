package routers

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

func LoadArticle(e *gin.Engine)  {
	ArticleController := &controller.ArticleController{}
	Article := e.Group("/api/v1/article")
	{
		// 创建文章
		Article.POST("create",ArticleController.CreateArticle)
		// 查询文章
		Article.GET("all",ArticleController.GetArticle)
	}
}