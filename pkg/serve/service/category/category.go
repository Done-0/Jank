// Package service 提供业务逻辑处理，处理类目相关业务
// 创建者：Done-0
// 创建时间：2025-05-10
package service

import (
	"fmt"
	"strings"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/category"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/controller/category/dto"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/category"
)

// GetCategoryByID 根据 ID 获取类目
// 参数：
//   - c: Echo 上下文
//   - req: 获取类目请求
//
// 返回值：
//   - interface{}: 获取到的类目视图对象
//   - error: 操作过程中的错误
func GetCategoryByID(c echo.Context, req *dto.GetOneCategoryRequest) (*category.CategoriesVO, error) {
	cat, err := mapper.GetCategoryByID(c, req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("根据 ID 获取类目失败: %v", err)
		return nil, fmt.Errorf("根据 ID 获取类目失败: %w", err)
	}

	categoryVO, err := utils.MapModelToVO(cat, &category.CategoriesVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("获取类目时映射 VO 失败: %w", err)
	}

	return categoryVO.(*category.CategoriesVO), nil
}

// GetCategoryTree 获取类目树
// 参数：
//   - c: Echo 上下文
//
// 返回值：
//   - []*category.CategoriesVO: 类目树结构
//   - error: 操作过程中的错误
func GetCategoryTree(c echo.Context) ([]*category.CategoriesVO, error) {
	categories, err := mapper.GetAllActivatedCategories(c)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目树失败: %v", err)
		return nil, fmt.Errorf("获取类目树失败: %w", err)
	}

	// 构建类目映射（ID -> 类目对象指针）
	categoryMap := make(map[int64]*model.Category)
	for i := range categories {
		categoryMap[categories[i].ID] = categories[i]
	}

	// 构建父子关系
	var rootCategories []*model.Category
	for i := range categories {
		cat := categories[i]
		if cat.ParentID == 0 {
			rootCategories = append(rootCategories, cat)
		} else if parent, exists := categoryMap[cat.ParentID]; exists {
			if parent.Children == nil {
				parent.Children = make([]*model.Category, 0)
			}
			parent.Children = append(parent.Children, cat)
		}
	}

	var rootCategoriesVO []*category.CategoriesVO
	for _, root := range rootCategories {
		rootCategoryVO, err := buildCategoryVOTree(c, root)
		if err != nil {
			utils.BizLogger(c).Errorf("获取类目树时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取类目树时映射 VO 失败: %w", err)
		}
		rootCategoriesVO = append(rootCategoriesVO, rootCategoryVO)
	}

	return rootCategoriesVO, nil
}

// GetCategoryChildrenByID 根据类目 ID 获取层级子类目
// 参数：
//   - c: Echo 上下文
//   - req: 获取类目请求
//
// 返回值：
//   - []*category.CategoriesVO: 子类目列表
//   - error: 操作过程中的错误
func GetCategoryChildrenByID(c echo.Context, req *dto.GetOneCategoryRequest) ([]*category.CategoriesVO, error) {
	categories, err := mapper.GetAllActivatedCategories(c)
	if err != nil {
		utils.BizLogger(c).Errorf("根据 ID 获取层级子类目失败: %v", err)
		return nil, fmt.Errorf("根据 ID 获取层级子类目失败: %w", err)
	}

	// 构建类目映射与父子关系
	categoryMap := make(map[int64]*model.Category)
	for i := range categories {
		categoryMap[categories[i].ID] = categories[i]
	}

	for i := range categories {
		cat := categories[i]
		if cat.ParentID != 0 {
			if parent, exists := categoryMap[cat.ParentID]; exists {
				if parent.Children == nil {
					parent.Children = make([]*model.Category, 0)
				}
				parent.Children = append(parent.Children, cat)
			}
		}
	}

	// 查找目标类目
	target, exists := categoryMap[req.ID]
	if !exists {
		utils.BizLogger(c).Errorf("类目 ID=%d 不存在", req.ID)
		return nil, fmt.Errorf("类目 ID=%d 不存在", req.ID)
	}

	var childrenVO []*category.CategoriesVO
	if len(target.Children) > 0 {
		for _, child := range target.Children {
			childVO, err := buildCategoryVOTree(c, child)
			if err != nil {
				utils.BizLogger(c).Errorf("获取层级子类目时映射 VO 失败: %v", err)
				return nil, fmt.Errorf("获取层级子类目时映射 VO 失败: %w", err)
			}
			childrenVO = append(childrenVO, childVO)
		}
	}

	return childrenVO, nil
}

// CreateCategory 创建类目
// 参数：
//   - c: Echo 上下文
//   - req: 创建类目请求
//
// 返回值：
//   - *category.CategoriesVO: 创建后的类目视图对象
//   - error: 操作过程中的错误
func CreateCategory(c echo.Context, req *dto.CreateOneCategoryRequest) (*category.CategoriesVO, error) {
	var categoryVO *category.CategoriesVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		newCategory := &model.Category{
			Name:        req.Name,
			Description: req.Description,
			ParentID:    req.ParentID,
			Path:        "",
		}

		// 处理父类目路径
		if req.ParentID != 0 {
			parentCat, err := mapper.GetCategoryByID(c, req.ParentID)
			if err != nil {
				utils.BizLogger(c).Errorf("获取父类目失败: %v", err)
				return fmt.Errorf("获取父类目失败: %w", err)
			}

			if parentCat.Path == "" {
				newCategory.Path = fmt.Sprintf("/%d", req.ParentID)
			} else {
				newCategory.Path = fmt.Sprintf("%s/%d", parentCat.Path, req.ParentID)
			}
		}

		// 创建类目
		if err := mapper.CreateCategory(c, newCategory); err != nil {
			utils.BizLogger(c).Errorf("创建类目失败: %v", err)
			return fmt.Errorf("创建类目失败: %w", err)
		}

		vo, err := utils.MapModelToVO(newCategory, &category.CategoriesVO{})
		if err != nil {
			utils.BizLogger(c).Errorf("创建类目时映射 VO 失败: %v", err)
			return fmt.Errorf("创建类目时映射 VO 失败: %w", err)
		}

		categoryVO = vo.(*category.CategoriesVO)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return categoryVO, nil
}

// UpdateCategory 更新类目
// 参数：
//   - c: Echo 上下文
//   - req: 更新类目请求
//
// 返回值：
//   - *category.CategoriesVO: 更新后的类目视图对象
//   - error: 操作过程中的错误
func UpdateCategory(c echo.Context, req *dto.UpdateOneCategoryRequest) (*category.CategoriesVO, error) {
	var updatedVO *category.CategoriesVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		existingCategory, err := mapper.GetCategoryByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取类目失败: %v", err)
			return fmt.Errorf("获取类目失败: %w", err)
		}

		if req.ParentID == req.ID {
			utils.BizLogger(c).Error("父类目不能设置为自身")
			return fmt.Errorf("父类目不能设置为自身")
		}

		var parentPath string
		if req.ParentID != 0 {
			parentCategory, err := mapper.GetCategoryByID(c, req.ParentID)
			if err != nil {
				utils.BizLogger(c).Errorf("获取父类目「%d」失败: %v", req.ParentID, err)
				return fmt.Errorf("获取父类目「%d」失败: %w", req.ParentID, err)
			}

			if parentCategory.Path != "" && strings.Contains(parentCategory.Path, fmt.Sprintf("/%d/", req.ID)) {
				utils.BizLogger(c).Errorf("检测到循环引用: 无法将子类目「%d」设置为父类目", req.ID)
				return fmt.Errorf("检测到循环引用: 无法将子类目「%d」设置为父类目", req.ID)
			}

			parentPath = parentCategory.Path
		}

		oldPath := existingCategory.Path
		existingCategory.Name = req.Name
		existingCategory.Description = req.Description
		existingCategory.ParentID = req.ParentID

		if req.ParentID == 0 {
			existingCategory.Path = ""
		} else if parentPath == "" {
			existingCategory.Path = fmt.Sprintf("/%d", req.ParentID)
		} else {
			existingCategory.Path = fmt.Sprintf("%s/%d", parentPath, req.ParentID)
		}

		if err := mapper.UpdateCategory(c, existingCategory); err != nil {
			utils.BizLogger(c).Errorf("更新类目失败: %v", err)
			return fmt.Errorf("更新类目失败: %w", err)
		}

		if oldPath != existingCategory.Path {
			if err := recursivelyUpdateChildrenPaths(c, existingCategory); err != nil {
				utils.BizLogger(c).Errorf("更新子类目路径失败: %v", err)
				return fmt.Errorf("更新子类目路径失败: %w", err)
			}
		}

		vo, err := buildCategoryVOTree(c, existingCategory)
		if err != nil {
			utils.BizLogger(c).Errorf("更新类目时映射 VO 失败: %v", err)
			return fmt.Errorf("更新类目时映射 VO 失败: %w", err)
		}

		updatedVO = vo
		return nil
	})

	if err != nil {
		return nil, err
	}

	return updatedVO, nil
}

// DeleteCategory 软删除类目
// 参数：
//   - c: Echo 上下文
//   - req: 删除类目请求
//
// 返回值：
//   - []*category.CategoriesVO: 被删除的类目树结构
//   - error: 操作过程中的错误
func DeleteCategory(c echo.Context, req *dto.DeleteOneCategoryRequest) ([]*category.CategoriesVO, error) {
	var deletedCategoriesVO []*category.CategoriesVO

	err := utils.RunDBTransaction(c, func(tx error) error {
		cat, err := mapper.GetCategoryByID(c, req.ID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取类目失败: %v", err)
			return fmt.Errorf("获取类目失败: %w", err)
		}

		categoriesToDelete, err := mapper.GetCategoriesByPath(c, cat.Path)
		if err != nil {
			utils.BizLogger(c).Errorf("获取子类目失败: %v", err)
			return fmt.Errorf("获取子类目失败: %w", err)
		}

		// 删除对应类目及其子类目并返回被删除的类目树
		var deletedCategories []*model.Category
		var categoryIDs []int64

		categoryIDs = append(categoryIDs, req.ID)
		deletedCategories = append(deletedCategories, cat)

		for _, childCat := range categoriesToDelete {
			if childCat.ID != req.ID && strings.HasPrefix(childCat.Path, cat.Path) {
				deletedCategories = append(deletedCategories, childCat)
				categoryIDs = append(categoryIDs, childCat.ID)
			}
		}

		// 删除相关文章-类目关联
		for _, categoryID := range categoryIDs {
			if err := mapper.DeletePostCategoryByCategoryID(c, categoryID); err != nil {
				utils.BizLogger(c).Errorf("删除类目「%d」的文章关联失败: %v", categoryID, err)
				return fmt.Errorf("删除类目「%d」的文章关联失败: %w", categoryID, err)
			}
		}

		// 软删除类目
		if err := mapper.DeleteCategoriesByPathSoftly(c, cat.Path, req.ID); err != nil {
			utils.BizLogger(c).Errorf("软删除类目失败: %v", err)
			return fmt.Errorf("软删除类目失败: %w", err)
		}

		// 构建被删除类目的树形结构
		deletedCategoryMap := make(map[int64]*model.Category)
		for i := range deletedCategories {
			deletedCategoryMap[deletedCategories[i].ID] = deletedCategories[i]
		}

		var rootDeletedCategories []*model.Category
		for i := range deletedCategories {
			cat := deletedCategories[i]
			if cat.ParentID == 0 {
				// 如果没有父类目，直接作为根节点
				rootDeletedCategories = append(rootDeletedCategories, cat)
			} else if _, exists := deletedCategoryMap[cat.ParentID]; !exists {
				// 如果父类目不在被删除的类目中，作为根节点
				rootDeletedCategories = append(rootDeletedCategories, cat)
			} else if parent, exists := deletedCategoryMap[cat.ParentID]; exists {
				// 父类目存在，添加为其子节点
				if parent.Children == nil {
					parent.Children = make([]*model.Category, 0)
				}
				parent.Children = append(parent.Children, cat)
			}
		}

		for _, rootCat := range rootDeletedCategories {
			rootVO, err := buildCategoryVOTree(c, rootCat)
			if err != nil {
				utils.BizLogger(c).Errorf("构建删除类目树 VO 失败: %v", err)
				return fmt.Errorf("构建删除类目树 VO 失败: %w", err)
			}
			deletedCategoriesVO = append(deletedCategoriesVO, rootVO)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return deletedCategoriesVO, nil
}

// recursivelyUpdateChildrenPaths 递归更新子类目路径
// 参数：
//   - c: Echo 上下文
//   - parentCategory: 父类目
//
// 返回值：
//   - error: 操作过程中的错误
func recursivelyUpdateChildrenPaths(c echo.Context, parentCategory *model.Category) error {
	children, err := mapper.GetCategoriesByParentID(c, parentCategory.ID)
	if err != nil {
		return fmt.Errorf("获取子类目失败: %w", err)
	}

	for _, child := range children {
		if parentCategory.Path == "" {
			child.Path = fmt.Sprintf("/%d", parentCategory.ID)
		} else {
			child.Path = fmt.Sprintf("%s/%d", parentCategory.Path, parentCategory.ID)
		}

		if err := mapper.UpdateCategory(c, child); err != nil {
			return fmt.Errorf("更新子类目路径失败: %w", err)
		}

		if err := recursivelyUpdateChildrenPaths(c, child); err != nil {
			return err
		}
	}

	return nil
}

// buildCategoryVOTree 构建类目树 VO
// 参数：
//   - c: Echo 上下文
//   - cat: 类目对象
//
// 返回值：
//   - *category.CategoriesVO: 构建后的类目树视图对象
//   - error: 操作过程中的错误
func buildCategoryVOTree(c echo.Context, cat *model.Category) (*category.CategoriesVO, error) {
	categoryVO, err := utils.MapModelToVO(cat, &category.CategoriesVO{})
	if err != nil {
		return nil, fmt.Errorf("构建类目树 VO 失败: %w", err)
	}

	catVO := categoryVO.(*category.CategoriesVO)
	catVO.Children = make([]*category.CategoriesVO, 0)

	if len(cat.Children) > 0 {
		for _, child := range cat.Children {
			childVO, err := buildCategoryVOTree(c, child)
			if err != nil {
				return nil, err
			}
			catVO.Children = append(catVO.Children, childVO)
		}
	}

	return catVO, nil
}
