package catagory

import (
	"blog/database"
	"blog/models"
	"fmt"
	"testing"
	"time"
)

func init() {
	err := database.InitMySql()
	if err != nil {
		return
	}
}

// 调用创建分类接口
func TestAddCategory(t *testing.T) {
	category := &models.ArticleCategory{
		CategoryName: "C++",
		CreateTime:   time.Now(),
		UpdateTime:   time.Now(),
	}
	//database.DB.Debug().Create(category)
	err := AddCategory(category)
	if err != nil{
		// 创建失败
		fmt.Println(err)
		return
	}

}
// 查询所有分类
func TestGetAllCategory(t *testing.T) {
	all, err := GetAllCategory()
	if err != nil{
		t.Logf("error:%+v", err)
		return
	}
	fmt.Println(all)
}

// 删除
func TestDeleteCategory(t *testing.T) {
	err := DeleteCategory(2)

	if err!=nil{
		return
	}
}