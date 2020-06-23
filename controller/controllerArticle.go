package controller

import (
	"blog/dao/articleTag"
	"blog/dao/articles"
	"blog/models"
	"blog/dao/tags"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type ArticleController struct {
}

// 创建文章
func (a *ArticleController) CreateArticle(c *gin.Context) {
	title := c.PostForm("title")
	categoryid := c.PostForm("category")
	ossURL := c.PostForm("ossUrl")
	//status := c.PostForm("status")
	html := c.PostForm("html")
	md := c.PostForm("md")
	tag := c.PostForm("tags")
	tagsSplit := strings.Split(tag, ";")
	fmt.Println(tagsSplit)

	fmt.Printf("categoryid:%+v\n", categoryid)
	category, err := strconv.Atoi(categoryid)
	//if status == "true" {
	//	statu := true
	//} else {
	//	statu := false
	//}

	article := &models.Article{
		Title:        title,
		CreateTime:   time.Now().UnixNano() / 1e6,
		UpdateTime:   time.Now().UnixNano() / 1e6,
		Status:       true,
		Md:           md,
		Html:         html,
		CoverAddress: ossURL,
		Author:       "Lz12Code",
		Top:          0,
		CategoryId:   category,
		Summary:      "阿斯顿撒",
		Views:        888,
	}
	// 调用创建文章函数
	err = articles.CreateAticle(article)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}

	for _,tag := range tagsSplit {

		// 调用添加标签函数
		tag := &models.Tag{
			TagName:    tag,
			CreateTime: time.Now().UnixNano() / 1e6,
			UpdateTime: time.Now().UnixNano() / 1e6,
		}
		err = tags.CreateTag(tag)
		if err != nil {
			return
		}

		// 调用给第三张表 文章标签表添加记录的函数
		articleTags := &models.ArticleTag{
			ArticleId:  article.Id,
			TagId:      tag.Id,
			CreateTime: time.Now().UnixNano() / 1e6,
			UpdateTime: time.Now().UnixNano() / 1e6,
		}
		err = articleTag.CreateArticleTag(articleTags)
		if err != nil {
			return
		}
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
