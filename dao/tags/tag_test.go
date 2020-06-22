package tags

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

func TestCreateTag(t *testing.T) {
	category := &models.ArticleTag{
		TagName:    "eeee",
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6,
	}

	err := CreateTag(category)
	if err != nil {
		// 创建失败
		fmt.Println(err)
		return
	}
}

func TestQueryTagName(t *testing.T) {
	a := make(map[int][]int)
	a[5] = []int{5, 6}
	a[4] = []int{2, 3}
	_, err := QueryTagName(a)
	if err != nil {
		return
	}
}
