package articles

import (
	"blog/database"
	"blog/models"
)

// 发布文章
func CreateAticle(article *models.ArticleList) (err error) {
	err = database.DB.Debug().Create(&article).Error
	if err != nil {
		return
	}
	return
}


// 文章列表
func GetArticle() (articleList []models.ArticleList,err error) {
	// 根据update_time排序 倒序
	err = database.DB.Debug().Order("update_time desc").Find(&articleList).Error
	if err != nil {
		return
	}

	//categoryList := make([]int,0)
	//for _, v := range articleList {
	//	categoryList = append(categoryList, v.CategoryId)
	//}
	//
	//categoryMap, err := catagory.QueryCategoryName(categoryList)
	//if err != nil {
	//	return articleList, err
	//}
	//fmt.Printf("categoryMap:%+v", categoryMap)

	return
}