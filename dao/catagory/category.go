package catagory

import (
	"blog/database"
	"blog/models"
)

// 创建分类
func CreateCategory(category *models.Category) (err error) {
	err = database.DB.Debug().Create(&category).Error
	if err != nil {
		return
	}
	return
}

// 查询分类 查询所有 返回切片类型
func GetAllCategory() (all []models.Category, err error) {
	// 根据update_time排序 倒序
	err = database.DB.Debug().Order("update_time desc").Find(&all).Error
	if err != nil {
		return
	}
	return
}

// 修改分类
func ModifyCategory(category *models.Category) (err error) {
	err = database.DB.Debug().Model(&models.Category{}).Where("id = ?", category.Id).Update(category).Error
	if err != nil {
		return
	}
	return
}

// 删除单个分类
func DeleteCategory(categoryId int) (err error) {

	err = database.DB.Debug().Where("id = ?", categoryId).Delete(&models.Category{}).Error
	if err != nil {
		return
	}
	return
}

// 查询所有文章的对应的所有分类  连表查询
func QueryCategoryName(categorySlice []int) (map[int]string, error) {
	category := make([]models.Category, 0)
	categoryMap := make(map[int]string)

	if err := database.DB.Debug().Where("id in (?)", categorySlice).Find(&category).Error; err != nil {
		return categoryMap, err
	}

	for _, v := range category {
		categoryMap[v.Id] = v.CategoryName
	}

	return categoryMap, nil
}


