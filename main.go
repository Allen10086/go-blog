package main

import (
	"blog/routers"

)

func main() {
	// 路由
	r := routers.SetupRouter()
	r.Run(":8081")

}
