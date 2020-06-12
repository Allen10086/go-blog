package routers

import (
	"blog/controller"
	"github.com/gin-gonic/gin"
)

// 文章分类路由
func LoadCategory(e *gin.Engine) {
	CategoryControl := &controller.CategoryController{}
	// 文章分类路由
	CategoryG := e.Group("/api/v1/category/")
	{   // 查询all
		CategoryG.GET("/all", CategoryControl.GetCategory)
		// 添加
		CategoryG.POST("/create", CategoryControl.CreateCategory)
		// 删除
		CategoryG.DELETE("/delete",CategoryControl.DeleteCategory)
		// 修改
		CategoryG.PUT("/modify",CategoryControl.ModifyCategory)

	}
}
