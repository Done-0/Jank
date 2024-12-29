package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	model "jank.com/jank_blog/internal/model/category"
)

const (
	activeCategoryCondition = "`is_active` = true"
)

// GetCategoryByID 根据 ID 查找类目
func GetCategoryByID(id int64) (*model.Category, error) {
	var category model.Category
	err := global.DB.Where("id = ? AND is_active = ?", id, true).First(&category).Error
	if err != nil {
		return nil, err
	}
	return &category, nil
}

// GetCategoriesByParentID 根据父类目 ID 查找直接子类目
func GetCategoriesByParentID(parentID int64) ([]model.Category, error) {
	var categories []model.Category
	err := global.DB.Where("parent_id = ?", parentID).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoriesByParentPath 根据父类目的路径查找所有子类目
func GetCategoriesByParentPath(parentPath string) ([]model.Category, error) {
	var categories []model.Category
	err := global.DB.Where("path LIKE ?", fmt.Sprintf("%s%%", parentPath)).Find(&categories).Error
	if err != nil {
		return nil, err
	}
	return categories, nil
}

// GetCategoriesByPath 根据路径获取所有子类目
func GetCategoriesByPath(path string) ([]*model.Category, error) {
	var categories []*model.Category
	err := global.DB.Model(&model.Category{}).
		Where("path LIKE ?", fmt.Sprintf("%s%%", path)).
		Find(&categories).Error
	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetAllCategories 获取所有有效类目
func GetAllActivedCategories() ([]*model.Category, error) {
	var categories []*model.Category
	err := global.DB.Where(activeCategoryCondition).
		Find(&categories).Error

	if err != nil {
		return nil, err
	}

	return categories, nil
}

// GetParentCategoryPathByID 根据父类目 ID 查找父类目的路径
func GetParentCategoryPathByID(parentID int64) (string, error) {
	if parentID == 0 {
		return "", nil
	}
	var parentCategory *model.Category
	err := global.DB.Select("path").Where("id = ? AND is_active = ?", parentID, true).First(&parentCategory).Error
	if err != nil {
		return "", err
	}
	return parentCategory.Path, nil
}

// CreateCategory 将新类目保存到数据库
func CreateCategory(newCategory *model.Category) error {
	return global.DB.Create(newCategory).Error
}

// UpdateCategory 更新类目信息
func UpdateCategory(category *model.Category) error {
	return global.DB.Save(category).Error
}

// DeleteCategoriesByPathSoftly 软删除类目及其子类目
func DeleteCategoriesByPathSoftly(path string) error {
	return global.DB.Model(&model.Category{}).
		Where("path LIKE ?", fmt.Sprintf("%s%%", path)).
		Update("is_active", false).Error
}
