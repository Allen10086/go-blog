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

type Tags struct {
	TagName string
}

// 查询所有文章对应的所有标签    联表查询   GORM查询出来的值只能用结构体接
func QueryTagName(tagMap map[int][]int) (map[int][]string, error) {
	tag := make([]Tags, 0)
	tags := make(map[int][]string)
	// 查询所有文章的所有分类 添加到tag切片内
	for i, v := range tagMap {
		t := make([]string, 0)
		if err := database.DB.Debug().Select("tag_name").Table("article_tags").Where("id in (?)", v).Find(&tag).Error; err != nil {
			return tags, err
		}
		for _, v := range tag {
			t = append(t, v.TagName)
		}
		tags[i] = t
	}
	return tags, nil
}
