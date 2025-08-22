// Package mapper 提供分类相关的数据访问接口
// 创建者：Done-0
// 创建时间：2025-08-13
package mapper

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/category"
)

// CategoryMapper 分类数据访问接口
type CategoryMapper interface {
	GetCategoryByID(c *app.RequestContext, categoryID int64) (*category.Category, error)                                                   // 根据 ID 获取分类
	ListCategories(c *app.RequestContext, pageNo, pageSize int64, parentID *int64, isActive *bool) ([]*category.Category, int64, error) // 获取分类列表，支持按父分类和状态筛选
	CreateCategory(c *app.RequestContext, category *category.Category) error                                                             // 创建分类
	UpdateCategory(c *app.RequestContext, category *category.Category) error                                                             // 更新分类
	DeleteCategory(c *app.RequestContext, categoryID int64) error                                                                        // 删除分类
}
