package controller

import (
	"blog/dao/articles"
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type ArticleController struct {
}

// 创建文章
func (a *ArticleController) CreateArticle(c *gin.Context) {
	article := &models.Article{
		Title:        "博客开发",
		CreateTime:   time.Now().UnixNano() / 1e6,
		UpdateTime:   time.Now().UnixNano() / 1e6,
		Status:       true,
		Md:           "md",
		Html:         "html",
		CoverAddress: "https://lz12code.oss-cn-beijing.aliyuncs.com/myblog/153f73d9-1635-4e64-98c3-025cb3d04b43.jpg",
		Author:       "Lz12Code",
		Top:          0,
		CategoryId:   39,
		Summary:      "阿斯顿撒",
		Views:        888,
	}
	err := articles.CreateAticle(article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})

}

// 查询文章
func (a *ArticleController) GetArticle(c *gin.Context) {
	articleList, err := articles.GetArticle()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 失败返回code 1001
			"code":    1001,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":        1000,
		"message":     "success",
		"articleList": articleList,
	})
	return
}
