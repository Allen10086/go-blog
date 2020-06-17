package models

type UserInfo struct {
	Id       int    `json:"id" gorm:"not null pk autoincr TINYINT(4) 'id'"`
	UserName string `json:"user_name" gorm:"default 'NULL' VARCHAR(255) 'user_name'"`
	PassWord string `json:"pass_word" gorm:"default 'NULL' VARCHAR(255) 'pass_word'"`
}