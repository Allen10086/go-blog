package articles

import (
	"blog/dao/articleTag"
	"blog/dao/catagory"
	"blog/dao/tags"
	"blog/database"
	"blog/models"
	"fmt"
)

// 发布文章
func CreateAticle(article *models.Article) (err error) {
	err = database.DB.Debug().Create(&article).Error
	if err != nil {
		return
	}
	return
}

// 文章管理列表
type ArticleList struct {
	ID           int      `json:"id"`
	Title        string   `json:"title"`
	CategoryName string   `json:"category_name"`
	CoverAddress string   `json:"cover_address"`
	CreateTime   int64    `json:"create_time"`
	UpdateTime   int64    `json:"update_time"`
	Views        int      `json:"views"`
	Status       bool     `json:"status"`
	TagName      []string `json:"tag_name"`
	//models.Tag
	//models.Article
}

// 文章管理列表
func GetArticle() (articleList []ArticleList, err error) {
	articles := make([]models.Article, 0)
	// 根据update_time排序 倒序
	err = database.DB.Debug().Order("update_time desc").Find(&articles).Error
	if err != nil {
		return
	}

	// 分类id列表
	categoryList := make([]int, 0)
	// 文章id列表
	articleIdList := make([]int, 0)

	for _, v := range articles {
		categoryList = append(categoryList, v.CategoryId)
		articleIdList = append(articleIdList, v.Id)
	}

	// 分类查询
	categoryMap, err := catagory.QueryCategoryName(categoryList)
	if err != nil {
		return articleList, err
	}

	// 文章标签对应关系查询
	articleTagList, err := articleTag.QueryArticleTag()
	if err != nil {
		return
	}
	fmt.Println(articleTagList)

	// 标签查询
	tagMap, err := tags.QueryTags(articleIdList)
	if err != nil {
		return
	}
	fmt.Printf("tagMap:%+v\n", tagMap)

	// 组装文章id -> 标签名称 数据结构

	list := make([]ArticleList, len(articles))
	if len(articles) > 0 {
		for i := range articles {
			list[i].ID = articles[i].Id
			list[i].Title = articles[i].Title
			list[i].CategoryName = categoryMap[articles[i].CategoryId]
			list[i].CoverAddress = articles[i].CoverAddress
			list[i].CreateTime = articles[i].CreateTime
			list[i].UpdateTime = articles[i].UpdateTime
			list[i].Views = articles[i].Views
			list[i].Status = articles[i].Status
			list[i].TagName = tagMap[articles[i].Id]
		}
	}

	//log.Printf("List:%+v", list)

	articleList = list
	return
}
