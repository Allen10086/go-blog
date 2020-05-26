package main

import (
	"blog/dao"
	"blog/routers"
)

func main() {
	// 数据库初始化
	err := dao.InitMySql()
	if err != nil {
		panic(err)
	}
	// 关闭连接
	defer dao.DB.Close()
	// 路由
	r := routers.SetupRouter()
	r.Run(":8081")

}
