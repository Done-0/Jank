// Package impl 分类服务实现
// 创建者：Done-0
// 创建时间：2025-08-13
package impl

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/category"
	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

// CategoryServiceImpl 分类服务实现
type CategoryServiceImpl struct {
	categoryMapper mapper.CategoryMapper
}

// NewCategoryService 创建分类服务实例
func NewCategoryService(categoryMapperImpl mapper.CategoryMapper) service.CategoryService {
	return &CategoryServiceImpl{
		categoryMapper: categoryMapperImpl,
	}
}

// GetCategory 获取单个分类
func (cs *CategoryServiceImpl) GetCategory(c *app.RequestContext, req *dto.GetCategoryRequest) (*vo.GetCategoryResponse, error) {
	categoryID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid category ID format: %s", req.ID)
		return nil, fmt.Errorf("invalid category ID format: %w", err)
	}

	category, err := cs.categoryMapper.GetCategoryByID(c, categoryID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get category with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	return &vo.GetCategoryResponse{
		ID:          strconv.FormatInt(category.ID, 10),
		Name:        category.Name,
		Description: category.Description,
		ParentID:    strconv.FormatInt(category.ParentID, 10),
		Sort:        category.Sort,
		IsActive:    category.IsActive,
		CreatedAt:   time.Unix(category.GmtCreated, 0).Format("2006-01-02 15:04:05"),
		UpdatedAt:   time.Unix(category.GmtModified, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

// ListCategories 获取分类列表
func (cs *CategoryServiceImpl) ListCategories(c *app.RequestContext, req *dto.ListCategoriesRequest) (*vo.ListCategoriesResponse, error) {
	var parentID *int64
	if req.ParentID != "" {
		pid, err := strconv.ParseInt(req.ParentID, 10, 64)
		if err != nil {
			logger.BizLogger(c).Errorf("invalid parent ID format: %s", req.ParentID)
			return nil, fmt.Errorf("invalid parent ID format: %w", err)
		}
		parentID = &pid
	}

	categories, total, err := cs.categoryMapper.ListCategories(c, req.PageNo, req.PageSize, parentID, req.IsActive)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list categories: %v", err)
		return nil, fmt.Errorf("failed to list categories: %w", err)
	}

	var categoryItems []*vo.CategoryItem
	for _, cat := range categories {
		categoryItems = append(categoryItems, &vo.CategoryItem{
			ID:          strconv.FormatInt(cat.ID, 10),
			Name:        cat.Name,
			Description: cat.Description,
			ParentID:    strconv.FormatInt(cat.ParentID, 10),
			Sort:        cat.Sort,
			IsActive:    cat.IsActive,
			CreatedAt:   time.Unix(cat.GmtCreated, 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:   time.Unix(cat.GmtModified, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return &vo.ListCategoriesResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     categoryItems,
	}, nil
}

// Create 创建分类
func (cs *CategoryServiceImpl) Create(c *app.RequestContext, req *dto.CreateCategoryRequest) (*vo.CreateCategoryResponse, error) {
	var parentID int64 = 0
	if req.ParentID != "" {
		pid, err := strconv.ParseInt(req.ParentID, 10, 64)
		if err != nil {
			logger.BizLogger(c).Errorf("invalid parent ID format: %s", req.ParentID)
			return nil, fmt.Errorf("invalid parent ID format: %w", err)
		}
		parentID = pid
	}

	sort := req.Sort
	if sort == 0 {
		sort = 100 // 默认排序权重
	}

	category := &category.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    parentID,
		Sort:        sort,
		IsActive:    req.IsActive,
	}

	if err := cs.categoryMapper.CreateCategory(c, category); err != nil {
		logger.BizLogger(c).Errorf("failed to create category '%s': %v", req.Name, err)
		return nil, fmt.Errorf("failed to create category: %w", err)
	}

	logger.BizLogger(c).Infof("category created successfully with ID: %d", category.ID)

	return &vo.CreateCategoryResponse{
		ID:          strconv.FormatInt(category.ID, 10),
		Name:        category.Name,
		Description: category.Description,
		ParentID:    strconv.FormatInt(category.ParentID, 10),
		Sort:        category.Sort,
		IsActive:    category.IsActive,
		Message:     "Category created successfully",
	}, nil
}

// Update 更新分类
func (cs *CategoryServiceImpl) Update(c *app.RequestContext, req *dto.UpdateCategoryRequest) (*vo.UpdateCategoryResponse, error) {
	categoryID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid category ID format: %s", req.ID)
		return nil, fmt.Errorf("invalid category ID format: %w", err)
	}

	// 获取现有分类
	existingCategory, err := cs.categoryMapper.GetCategoryByID(c, categoryID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get category with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to get category: %w", err)
	}

	if req.Name != "" {
		existingCategory.Name = req.Name
	}
	existingCategory.Description = req.Description
	if req.ParentID != "" {
		parentID, err := strconv.ParseInt(req.ParentID, 10, 64)
		if err != nil {
			logger.BizLogger(c).Errorf("invalid parent ID format: %s", req.ParentID)
			return nil, fmt.Errorf("invalid parent ID format: %w", err)
		}
		existingCategory.ParentID = parentID
	}
	existingCategory.Sort = req.Sort
	existingCategory.IsActive = req.IsActive

	if err := cs.categoryMapper.UpdateCategory(c, existingCategory); err != nil {
		logger.BizLogger(c).Errorf("failed to update category with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to update category: %w", err)
	}

	logger.BizLogger(c).Infof("category updated successfully with ID: %d", existingCategory.ID)

	return &vo.UpdateCategoryResponse{
		ID:          strconv.FormatInt(existingCategory.ID, 10),
		Name:        existingCategory.Name,
		Description: existingCategory.Description,
		ParentID:    strconv.FormatInt(existingCategory.ParentID, 10),
		Sort:        existingCategory.Sort,
		IsActive:    existingCategory.IsActive,
		Message:     "Category updated successfully",
	}, nil
}

// Delete 删除分类
func (cs *CategoryServiceImpl) Delete(c *app.RequestContext, req *dto.DeleteCategoryRequest) (*vo.DeleteCategoryResponse, error) {
	categoryID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid category ID format: %s", req.ID)
		return nil, fmt.Errorf("invalid category ID format: %w", err)
	}

	if err := cs.categoryMapper.DeleteCategory(c, categoryID); err != nil {
		logger.BizLogger(c).Errorf("failed to delete category with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete category: %w", err)
	}

	logger.BizLogger(c).Infof("category deleted successfully with ID: %d", categoryID)

	return &vo.DeleteCategoryResponse{
		Message: "Category deleted successfully",
	}, nil
}
