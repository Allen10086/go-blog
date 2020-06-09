package models

// Todo 用户信息表
type UserInfos struct {
	ID       int8   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}