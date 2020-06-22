package articles

import (
	"blog/dao/catagory"
	"blog/dao/tags"
	"blog/database"
	"blog/models"
	"encoding/json"
	"log"
)

// 发布文章
func CreateAticle(article *models.ArticleList) (err error) {
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


}





// 文章列表
func GetArticle() (articleList []ArticleList, err error) {
	articles := make([]models.ArticleList, 0)
	// 根据update_time排序 倒序
	err = database.DB.Debug().Order("update_time desc").Find(&articles).Error
	if err != nil {
		return
	}
	//log.Printf("articles:%+v", articles)

	categoryList := make([]int, 0)

	//  tagsMap:map[int][]int{4:[]int{2, 3}, 5:[]int{5, 6}}  文章id对应的标签id
	tagsMap := make(map[int][]int)

	for _, v := range articles {
		tagId := make([]int, 0)
		err = json.Unmarshal([]byte(v.TagId), &tagId)
		if err != nil {
			return
		}
		//fmt.Printf("tagid:%#v", tagId)
		tagsMap[v.Id] = tagId
		// 将所有文章的分类id添加到切片内
		categoryList = append(categoryList, v.CategoryId)
	}
	//log.Printf("tagsMap:%#v",tagsMap)
	// 分类查询
	categoryMap, err := catagory.QueryCategoryName(categoryList)
	if err != nil {
		return articleList, err
	}
	//fmt.Printf("categoryMap:%+v", categoryMap)
	// 标签查询
	tagNameMap, err := tags.QueryTagName(tagsMap)
	if err != nil {
		return
	}

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
			list[i].TagName = tagNameMap[articles[i].Id]
		}
	}

	log.Printf("List:%+v", list)

	articleList = list
	return
}
