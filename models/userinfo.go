package models

type User struct {
	Id       int    `json:"id" gorm:"not null pk autoincr TINYINT(4) 'id'"`
	UserName string `json:"user_name" gorm:"default 'NULL' VARCHAR(255) 'user_name'"`
	PassWord string `json:"pass_word" gorm:"default 'NULL' VARCHAR(255) 'pass_word'"`
	Status   bool   `json:"status" gorm:"not null comment('状态 0:禁用  1:启动')  'status'"`
}
