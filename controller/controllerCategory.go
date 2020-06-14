package controller

import (
	"blog/dao/catagory"
	"blog/models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type CategoryController struct {
}

//type Category struct {
//	CategoryName string `from:"category_name" json:"category_name"`
//}

// 查询方法
func (w *CategoryController) GetCategory(c *gin.Context) {

	all, err := catagory.GetAllCategory()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			// 失败返回code 1001
			"code":    1001,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":         1000,
		"message":      "success",
		"CategoryList": all,
	})
	return

}



// 添加方法
func (w *CategoryController) CreateCategory(c *gin.Context) {
	categoryName := c.PostForm("category_name")
	category := &models.ArticleCategory{
		CategoryName: categoryName,
		CreateTime:   time.Now().UnixNano()/1e6,
		UpdateTime:   time.Now().UnixNano()/1e6,
	}
	err := catagory.CreateCategory(category)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})
	return

}



// 修改方法
func (w *CategoryController) ModifyCategory(c *gin.Context) {
	categoryName := c.PostForm("category_name")
	categoryId := c.PostForm("id")
	// 字符串转换成int
	Id, err := strconv.Atoi(categoryId)
	if err != nil {

		return
	}

	category := &models.ArticleCategory{
		Id:           Id,
		CategoryName: categoryName,
		UpdateTime:   time.Now().UnixNano()/1e6,
	}

	e := catagory.ModifyCategory(category)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": e,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})
	return

}




// 删除方法
func (w *CategoryController) DeleteCategory(c *gin.Context) {
	categoryId := c.Query("id")

	id, err := strconv.Atoi(categoryId)
	if err != nil {

		return
	}

	e := catagory.DeleteCategory(id)
	if e != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    1001,
			"message": "failed",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    1000,
		"message": "success",
	})
	return
}
