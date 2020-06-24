package tags

import (
	"blog/database"
	"blog/models"
)

// 添加标签
func CreateTag(tag *models.Tag) (err error) {
	err = database.DB.Debug().Create(&tag).Error
	if err != nil {
		return
	}
	return
}
// 根据名字查询标签
func QueryAllTagList(tagName string) (tag models.Tag, err error) {
	//db := database.DB.Debug().Where("tag_name = ?",tagName).Last(&tag)
	database.DB.Debug().Where("tag_name = ?",tagName).Last(&tag)
	//t := db.Value.(*models.Tag)
	//fmt.Println("t:", t)
	//if db.Error != nil {
	//	fmt.Printf("error:%+v\n", err)
	//	return tag, err
	//}
	//fmt.Println("tag:", tag)
	//tag = *t
	//return tag, nil
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