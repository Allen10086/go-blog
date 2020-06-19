package articles

import (
	"blog/database"
	"blog/models"
	"encoding/json"
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
	tagsId := make([]int, 0)
	tagsId = append(tagsId, 2, 3)
	tags, err := json.Marshal(&tagsId)
	article := &models.ArticleList{
		Title:        "vue",
		CreateTime:   time.Now().UnixNano() / 1e6,
		UpdateTime:   time.Now().UnixNano() / 1e6,
		Status:       true,
		Md:           "md",
		Html:         "html",
		CoverAddress: "https://lz12code.oss-cn-beijing.aliyuncs.com/myblog/153f73d9-1635-4e64-98c3-025cb3d04b43.jpg",
		Author:       "Lz12Code",
		Top:          0,
		CategoryId:   3,
		Summary:      "Lz12Code",
		Views:        888,
		TagId:        string(tags),
	}
	err = CreateAticle(article)
	if err != nil {
		return
	}
}

// 查询所有文章
func TestGetArticle(t *testing.T) {
	article, err := GetArticle()
	if err != nil {
		return
	}
	fmt.Printf("article:%+v", article)

	aa := make([]int, 0)
	t.Logf("aa len:%+v", len(aa))

	bb := make([]int, 3)
	t.Logf("bb len:%+v", len(bb))

	return
}
