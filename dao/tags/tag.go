package tags

import (
	"blog/database"
	"blog/models"
)

// 添加标签
func CreateTag(tag *models.Tag) (err error) {
	//err = database.DB.Debug().First(&article).Error
	err = database.DB.Debug().Create(&tag).Error
	if err != nil {
		return
	}
	return
}
// 查询所有标签
func QueryAllTagList() (tags []models.Tag, err error) {
	// 根据update_time排序 倒序
	err = database.DB.Debug().Order("update_time desc").Find(&tags).Error
	if err != nil {
		return
	}
	return
}

type Tags struct {
	TagName string
}

// 文章id 对应的标签
func QueryTags(articleIds []int) (map[int][]string, error)  {
	tagMap := make(map[int][]string)
	for _, id := range articleIds {
		t := make([]Tags, 0)
		name := make([]string, 0)
		if err := database.DB.Debug().Select("tags.tag_name").Table("article_tags").Joins("left join tags on article_tags.tag_id = tags.id").
			Where("article_tags.article_id = ?", id).Find(&t).Error; err != nil {
			return tagMap, err
		}

		for _, v := range t {
			name = append(name, v.TagName)
		}

		tagMap[id] = name
	}

	return tagMap, nil
}