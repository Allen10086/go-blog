package models

// 标签表
type ArticleTag struct {
	Id         int    `json:"id"gorm:"not null pk autoincr comment('id') INT(5) 'id'"`
	TagName    string `json:"tag_name" gorm:"not null comment('文章标签') VARCHAR(255) 'tag_name'"`
	CreateTime int64  `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime int64  `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
