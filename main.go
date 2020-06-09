package main

import (
	"blog/database"
	"blog/routers"
)

func main() {
	// 数据库初始化
	err := database.InitMySql()
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer database.DB.Close()
	// 路由
	r := routers.SetupRouter()
	r.Run(":8081")

}
