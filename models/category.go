package models


// 分类表
type Category struct {
	Id           int    `json:"id" gorm:"not null unique pk INT(11) 'id'"`
	CategoryName string `json:"category_name" gorm:"default 'NULL' CHAR(20) 'category_name'"`
	CreateTime   int64  `json:"create_time" gorm:"default 'NULL' BIGINT 'create_time'"`
	UpdateTime   int64  `json:"update_time" gorm:"default 'NULL' BIGINT 'update_time'"`
}
