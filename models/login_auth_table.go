package models

import (
	"blog/dao"
	"log"
)

// Todo 用户信息表
type UserInfos struct {
	ID       int8   `json:"id"`
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}

// Todo 查询用户名和密码
func GetUserInfo(username, password string) (user *UserInfos, err error) {
	user = &UserInfos{} // 赋值给user
	// 查询语句 去数据库中查询用户名和密码
	err = dao.DB.Table("user_infos").Where("user_name = ? and pass_word = ?", username, password).First(user).Error
	if err != nil {
		return nil, err
	}
	log.Printf("user:%+v", user)
	return
}
