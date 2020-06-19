package tags

import (
	"blog/database"
	"blog/models"
)




// 添加分类
func CreateTag(article *models.ArticleTag) (err error) {
	err = database.DB.Debug().Create(&article).Error
	if err != nil {
		return
	}
	return
}