package articles

import (
	"blog/dao/articleTag"
	"blog/dao/tags"
	"blog/database"
	"blog/models"
	"fmt"
	"testing"
	"time"
)

// 数据库初始化
func init() {
	err := database.InitMySql()
	if err != nil {
		return
	}
}

// 发布文章
func TestCreateAticle(t *testing.T) {
	tagsNames := make([]string, 0)
	tagsNames = append(tagsNames, "mysql", "mongo","redis")
	//tags, err := json.Marshal(&tagsName)
	article := &models.Article{
		Title:        "wahaha",
		CreateTime:   time.Now().UnixNano() / 1e6,
		UpdateTime:   time.Now().UnixNano() / 1e6,
		Status:       true,
		Md:           "md",
		Html:         "html",
		CoverAddress: "https://lz12code.oss-cn-beijing.aliyuncs.com/myblog/153f73d9-1635-4e64-98c3-025cb3d04b43.jpg",
		Author:       "Lz12Code",
		Top:          0,
		CategoryId:   7,
		Summary:      "Lz12Code",
		Views:        888,
	}
	// 调用创建文章函数
	err := CreateAticle(article)
	if err != nil {
		return
	}
	for _,tag := range tagsNames {

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
}

// 查询所有文章
func TestGetArticle(t *testing.T) {
	article, err := GetArticle()
	if err != nil {
		return
	}
	fmt.Printf("article:%+v", article)
	return
}
