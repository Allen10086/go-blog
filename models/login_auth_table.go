package models

import (
	"blog/dao"
)

// Todo 用户信息表
type UserInfo struct {
	ID       int8   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// Todo 查询用户名和密码
func GetUserInfo(username string, password string) (userinfo *UserInfo, err error) {
	dao.DB.AutoMigrate(&UserInfo{})
	// 4、查询
	user := &UserInfo{}
	err = dao.DB.Debug().Where(&UserInfo{UserName: username,PassWord: password}).First(user).Error
	if err != nil {
		return nil, err
	}
	return
}
