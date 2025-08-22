// Package impl 提供分类相关的数据访问实现
// 创建者：Done-0
// 创建时间：2025-08-13
package impl

import (
	"fmt"

	"github.com/cloudwego/hertz/pkg/app"
	"gorm.io/gorm"

	"github.com/Done-0/jank/internal/model/category"
	"github.com/Done-0/jank/internal/model/post"
	"github.com/Done-0/jank/internal/utils/db"
	"github.com/Done-0/jank/pkg/serve/mapper"
)

// CategoryMapperImpl 分类数据访问实现
type CategoryMapperImpl struct{}

// NewCategoryMapper 创建分类数据访问实例
func NewCategoryMapper() mapper.CategoryMapper {
	return &CategoryMapperImpl{}
}

// GetCategoryByID 根据ID获取分类
func (m *CategoryMapperImpl) GetCategoryByID(c *app.RequestContext, categoryID int64) (*category.Category, error) {
	var cat category.Category
	err := db.GetDBFromContext(c).Where("id = ? AND deleted = ?", categoryID, false).First(&cat).Error
	if err != nil {
		return nil, err
	}
	return &cat, nil
}

// ListCategories 获取分类列表，支持按父分类和状态筛选
func (m *CategoryMapperImpl) ListCategories(c *app.RequestContext, pageNo, pageSize int64, parentID *int64, isActive *bool) ([]*category.Category, int64, error) {
	var categories []*category.Category
	var total int64

	// 构建查询条件
	query := db.GetDBFromContext(c).Model(&category.Category{}).Where("deleted = ?", false)

	// 按父分类筛选
	if parentID != nil {
		query = query.Where("parent_id = ?", *parentID)
	}

	// 按状态筛选
	if isActive != nil {
		query = query.Where("is_active = ?", *isActive)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.Order("sort DESC, id ASC").Offset(int(offset)).Limit(int(pageSize)).Find(&categories).Error; err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

// CreateCategory 创建分类
func (m *CategoryMapperImpl) CreateCategory(c *app.RequestContext, cat *category.Category) error {
	return db.GetDBFromContext(c).Create(cat).Error
}

// UpdateCategory 更新分类
func (m *CategoryMapperImpl) UpdateCategory(c *app.RequestContext, cat *category.Category) error {
	return db.GetDBFromContext(c).Save(cat).Error
}

// DeleteCategory 删除分类（软删除，级联删除所有子分类）
func (m *CategoryMapperImpl) DeleteCategory(c *app.RequestContext, categoryID int64) error {
	dbConn := db.GetDBFromContext(c)

	var allCategoryIDs []int64
	var currentLevelIDs []int64
	currentLevelIDs = append(currentLevelIDs, categoryID)

	for len(currentLevelIDs) > 0 {
		var nextLevelIDs []int64

		var childCategories []category.Category
		dbConn.Select("id").Where("parent_id IN ? AND deleted = ?", currentLevelIDs, false).Find(&childCategories)

		for _, child := range childCategories {
			allCategoryIDs = append(allCategoryIDs, child.ID)
			nextLevelIDs = append(nextLevelIDs, child.ID)
		}

		currentLevelIDs = nextLevelIDs
	}

	allCategoryIDs = append(allCategoryIDs, categoryID)

	return dbConn.Transaction(func(tx *gorm.DB) error {
		if err := tx.Model(&post.Post{}).
			Where("category_id IN ? AND deleted = ?", allCategoryIDs, false).
			Update("category_id", nil).Error; err != nil {
			return fmt.Errorf("failed to clear post category references: %w", err)
		}

		if err := tx.Model(&category.Category{}).
			Where("id IN ? AND deleted = ?", allCategoryIDs, false).
			Update("deleted", true).Error; err != nil {
			return fmt.Errorf("failed to delete categories: %w", err)
		}

		return nil
	})
}
