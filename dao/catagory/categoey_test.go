package catagory

import (
	"blog/database"
	"blog/models"
	"fmt"
	"testing"
	"time"
)
// 数据库初始化
func init() {
	err := database.InitMySql()
	if err != nil {
		return
	}
}

// 调用创建分类接口
func TestCreateCategory(t *testing.T) {
	category := &models.Category{
		CategoryName: "tytyt",
		CreateTime:   time.Now().UnixNano()/1e6,
		UpdateTime:   time.Now().UnixNano()/1e6,
	}

	err := CreateCategory(category)
	if err != nil {
		// 创建失败
		fmt.Println(err)
		return
	}

}

// 查询所有分类
func TestGetAllCategory(t *testing.T) {
	all, err := GetAllCategory()
	if err != nil {
		t.Logf("error:%+v", err)
		return
	}
	fmt.Println(all)
}

// 修改
func TestModifyCategory(t *testing.T) {
	category := &models.Category{
		Id:           4,
		CategoryName: "Django",
		UpdateTime:   time.Now().UnixNano()/1e6,
	}
	err := ModifyCategory(category)
	if err != nil {
		return
	}
	return

}

// 删除
func TestDeleteCategory(t *testing.T) {
	err := DeleteCategory(8)

	if err != nil {
		return
	}
}
