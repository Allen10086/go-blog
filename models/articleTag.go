package models

// 文章标签表
type ArticleTag struct {
	Id         int   `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	ArticleId  int   `json:"article_id" gorm:"not null INT(11) 'article_id'"`
	TagId      int   `json:"tag_id" gorm:"not null INT(11) 'tag_id'"`
	CreateTime int64 `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime int64 `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
