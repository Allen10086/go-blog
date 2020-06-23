package models

// 标签表

type Tag struct {
	Id         int    `json:"id" gorm:"not null pk autoincr INT(11) 'id'"`
	TagName    string `json:"tag_name" gorm:"not null VARCHAR(255) 'tag_name'"`
	CreateTime int64 `json:"create_time" gorm:"default NULL BIGINT(20) 'create_time'"`
	UpdateTime int64 `json:"update_time" gorm:"default NULL BIGINT(20) 'update_time'"`
}