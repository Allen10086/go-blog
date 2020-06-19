package catagory

import (
	"blog/database"
	"testing"
)

func init()  {
	err := database.InitMySql()
	if err != nil {
		return
	}
}

func TestQueryCategoryName(t *testing.T) {
	categorySlice := []int{2,3}

	categoryMap, err := QueryCategoryName(categorySlice)
	if err != nil {
		t.Errorf("error:%+v", err)
	}

	t.Logf("categoryMap:%+v", categoryMap)
}
