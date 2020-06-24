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
	category := &models.Tag{
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
	//_, err := QueryTagName(a)
	//if err != nil {
	//	return
	//}
}

func TestQueryTags(t *testing.T) {
	ids := []int{23}
	tagMap, err := QueryTags(ids)
	if err != nil {
		t.Errorf("error:%+v", err)
	}
	t.Logf("tagMap:%+v", tagMap)
}

func TestQueryAllTagList(t *testing.T) {
	tags := &models.Tag{
		TagName:    "mysql1",
		CreateTime: time.Now().UnixNano() / 1e6,
		UpdateTime: time.Now().UnixNano() / 1e6,
	}
	tagss, err := QueryAllTagList(tags.TagName)
	if err!=nil{
		fmt.Println("error:", err)
		return
	}
	fmt.Printf("tagss:%+v",tagss.Id)
}