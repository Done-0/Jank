// Package mapper 提供数据模型与数据库交互的映射层，处理类目相关数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	category "jank.com/jank_blog/internal/model/category"
	"jank.com/jank_blog/internal/utils"
)

// GetCategoryByID 根据 ID 查找类目
// 参数：
//   - c: Echo 上下文
//   - id: 类目 ID
//
// 返回值：
//   - *category.Category: 类目信息
//   - error: 操作过程中的错误
func GetCategoryByID(c echo.Context, id int64) (*category.Category, error) {
	var cat category.Category
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", id, false).First(&cat).Error; err != nil {
		return nil, fmt.Errorf("获取类目失败: %v", err)
	}
	return &cat, nil
}

// GetCategoriesByParentID 根据父类目 ID 查找直接子类目
// 参数：
//   - c: Echo 上下文
//   - parentID: 父类目 ID
//
// 返回值：
//   - []*category.Category: 子类目列表
//   - error: 操作过程中的错误
func GetCategoriesByParentID(c echo.Context, parentID int64) ([]*category.Category, error) {
	var categories []*category.Category
	db := utils.GetDBFromContext(c)
	if err := db.Where("parent_id = ? AND deleted = ?", parentID, false).Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("获取子类目失败: %v", err)
	}
	return categories, nil
}

// GetCategoriesByPath 根据路径获取所有子类目
// 参数：
//   - c: Echo 上下文
//   - path: 类目路径
//
// 返回值：
//   - []*category.Category: 子类目列表
//   - error: 操作过程中的错误
func GetCategoriesByPath(c echo.Context, path string) ([]*category.Category, error) {
	var categories []*category.Category
	db := utils.GetDBFromContext(c)

	// 如果路径为空，使用特殊查询条件只查询子类目
	if path == "" {
		if err := db.Model(&category.Category{}).
			Where("deleted = ?", false).
			Find(&categories).Error; err != nil {
			return nil, fmt.Errorf("获取路径下类目失败: %v", err)
		}
	} else {
		// 对于非空路径，确保只返回以该路径开头的类目
		if err := db.Model(&category.Category{}).
			Where("path LIKE ? AND deleted = ?", fmt.Sprintf("%s%%", path), false).
			Find(&categories).Error; err != nil {
			return nil, fmt.Errorf("获取路径下类目失败: %v", err)
		}
	}

	return categories, nil
}

// GetAllActivatedCategories 获取所有未删除的类目
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - []*category.Category: 类目列表
//   - error: 操作过程中的错误
func GetAllActivatedCategories(c echo.Context) ([]*category.Category, error) {
	var categories []*category.Category
	db := utils.GetDBFromContext(c)
	if err := db.Where("deleted = ?", false).
		Find(&categories).Error; err != nil {
		return nil, fmt.Errorf("获取所有类目失败: %v", err)
	}
	return categories, nil
}

// CreateCategory 将新类目保存到数据库
// 参数：
//   - c: Echo 上下文
//   - newCategory: 类目信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreateCategory(c echo.Context, newCategory *category.Category) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(newCategory).Error; err != nil {
		return fmt.Errorf("创建类目失败: %v", err)
	}
	return nil
}

// UpdateCategory 更新类目信息
// 参数：
//   - c: Echo 上下文
//   - category: 类目信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateCategory(c echo.Context, category *category.Category) error {
	db := utils.GetDBFromContext(c)
	if err := db.Save(category).Error; err != nil {
		return fmt.Errorf("更新类目失败: %v", err)
	}
	return nil
}

// DeleteCategoriesByPathSoftly 软删除类目及其子类目
// 参数：
//   - c: Echo 上下文
//   - path: 类目路径
//   - id: 类目 ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteCategoriesByPathSoftly(c echo.Context, path string, id int64) error {
	db := utils.GetDBFromContext(c)
	if err := db.Model(&category.Category{}).
		Where("id = ? AND deleted = ?", id, false).
		Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除当前类目失败: %v", err)
	}

	if err := db.Model(&category.Category{}).
		Where("path LIKE ? AND deleted = ? AND path != ?", fmt.Sprintf("%s%%", path), false, path).
		Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除子类目失败: %v", err)
	}
	return nil
}
