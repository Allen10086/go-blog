package user

import (
	"blog/database"
	"blog/models"
	"log"
)

func init() {

}


// Todo 查询用户名和密码
func GetUserInfo(username, password string) (user *models.UserInfo, err error) {
	user = &models.UserInfo{} // 赋值给user
	// 查询语句 去数据库中查询用户名和密码
	err = database.DB.Table("user_infos").Where("user_name = ? and pass_word = ?", username, password).First(user).Error
	if err != nil {
		return nil, err
	}
	log.Printf("user:%+v", user)
	return

}
