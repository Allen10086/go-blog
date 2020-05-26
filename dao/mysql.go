package dao

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	DB *gorm.DB
)
// 初始化连接
func InitMySql() (err error) {
	dsn := "root:root@tcp(127.0.0.1:3306)/spiders?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	
	if err := db.DB().Ping(); err != nil {
		panic(err)
	}

	DB = db
	return
}

func Close()  {
	DB.Close()
}