package dao

import "github.com/jinzhu/gorm"

var (
	DB *gorm.DB
)
// 初始化连接
func InitMySql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/db1?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return DB.DB().Ping()
}

func Close()  {
	DB.Close()
}