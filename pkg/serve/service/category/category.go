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
func GetCategoryByID(req *dto.GetOneCategoryRequest, c echo.Context) (interface{}, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
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
func GetCategoryTree(c echo.Context) ([]*category.CategoriesVO, error) {
	categories, err := mapper.GetAllActivatedCategories()
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
		rootCategoryVO, err := buildCategoryVOTree(root, c)
		if err != nil {
			utils.BizLogger(c).Errorf("获取类目树时映射 VO 失败: %v", err)
			return nil, fmt.Errorf("获取类目树时映射 VO 失败: %w", err)
		}
		rootCategoriesVO = append(rootCategoriesVO, rootCategoryVO)
	}

	return rootCategoriesVO, nil
}

// GetCategoryChildrenByID 根据类目 ID 获取层级子类目
func GetCategoryChildrenByID(req *dto.GetOneCategoryRequest, c echo.Context) ([]*category.CategoriesVO, error) {
	categories, err := mapper.GetAllActivatedCategories()
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
			childVO, err := buildCategoryVOTree(child, c)
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
func CreateCategory(req *dto.CreateOneCategoryRequest, c echo.Context) (*category.CategoriesVO, error) {
	newCategory := &model.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
		Path:        "",
	}

	// 处理父类目路径
	if req.ParentID != 0 {
		parentCat, err := mapper.GetCategoryByID(req.ParentID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取父类目失败: %v", err)
			return nil, fmt.Errorf("获取父类目失败: %w", err)
		}

		if parentCat.Path == "" {
			newCategory.Path = fmt.Sprintf("/%d", req.ParentID)
		} else {
			newCategory.Path = fmt.Sprintf("%s/%d", parentCat.Path, req.ParentID)
		}
	}

	// 创建类目
	if err := mapper.CreateCategory(newCategory); err != nil {
		utils.BizLogger(c).Errorf("创建类目失败: %v", err)
		return nil, fmt.Errorf("创建类目失败: %w", err)
	}

	categoryVO, err := utils.MapModelToVO(newCategory, &category.CategoriesVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建类目时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("创建类目时映射 VO 失败: %w", err)
	}

	return categoryVO.(*category.CategoriesVO), nil
}

// UpdateCategory 更新类目
func UpdateCategory(req *dto.UpdateOneCategoryRequest, c echo.Context) (*category.CategoriesVO, error) {
	existingCategory, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败: %v", err)
		return nil, fmt.Errorf("获取类目失败: %w", err)
	}

	if req.ParentID == req.ID {
		utils.BizLogger(c).Error("父类目不能设置为自身")
		return nil, fmt.Errorf("父类目不能设置为自身")
	}

	var parentPath string
	if req.ParentID != 0 {
		parentCategory, err := mapper.GetCategoryByID(req.ParentID)
		if err != nil {
			utils.BizLogger(c).Errorf("获取父类目「%d」失败: %v", req.ParentID, err)
			return nil, fmt.Errorf("获取父类目「%d」失败: %w", req.ParentID, err)
		}

		if parentCategory.Path != "" && strings.Contains(parentCategory.Path, fmt.Sprintf("/%d/", req.ID)) {
			utils.BizLogger(c).Errorf("检测到循环引用: 无法将子类目「%d」设置为父类目", req.ID)
			return nil, fmt.Errorf("检测到循环引用: 无法将子类目「%d」设置为父类目", req.ID)
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

	if err := mapper.UpdateCategory(existingCategory); err != nil {
		utils.BizLogger(c).Errorf("更新类目失败: %v", err)
		return nil, fmt.Errorf("更新类目失败: %w", err)
	}

	if oldPath != existingCategory.Path {
		if err := recursivelyUpdateChildrenPaths(existingCategory, c); err != nil {
			utils.BizLogger(c).Errorf("更新子类目路径失败: %v", err)
			return nil, fmt.Errorf("更新子类目路径失败: %w", err)
		}
	}

	updatedVO, err := buildCategoryVOTree(existingCategory, c)
	if err != nil {
		utils.BizLogger(c).Errorf("更新类目时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("更新类目时映射 VO 失败: %w", err)
	}

	return updatedVO, nil
}

// DeleteCategory 软删除类目
func DeleteCategory(req *dto.DeleteOneCategoryRequest, c echo.Context) ([]*category.CategoriesVO, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败: %v", err)
		return nil, fmt.Errorf("获取类目失败: %w", err)
	}

	categoriesToDelete, err := mapper.GetCategoriesByPath(cat.Path)
	if err != nil {
		utils.BizLogger(c).Errorf("获取子类目失败: %v", err)
		return nil, fmt.Errorf("获取子类目失败: %w", err)
	}

	// 删除对应类目及其子类目并返回被删除的类目树
	var deletedCategories []*model.Category
	var categoryIDs []int64

	categoryIDs = append(categoryIDs, req.ID)
	deletedCategories = append(deletedCategories, cat)

	for _, cat := range categoriesToDelete {
		if cat.ID != req.ID && strings.HasPrefix(cat.Path, cat.Path) {
			deletedCategories = append(deletedCategories, cat)
			categoryIDs = append(categoryIDs, cat.ID)
		}
	}

	for _, categoryID := range categoryIDs {
		if err := mapper.DeletePostCategoryByCategoryID(categoryID); err != nil {
			utils.BizLogger(c).Errorf("删除文章-类目关联失败: %v", err)
			return nil, fmt.Errorf("删除文章-类目关联失败: %w", err)
		}
	}

	if err := mapper.DeleteCategoriesByPathSoftly(cat.Path, req.ID); err != nil {
		utils.BizLogger(c).Errorf("软删除类目失败: %v", err)
		return nil, fmt.Errorf("软删除类目失败: %w", err)
	}

	categoryMap := make(map[int64]*model.Category)
	for _, c := range deletedCategories {
		c.Children = []*model.Category{}
		categoryMap[c.ID] = c
	}

	for _, c := range deletedCategories {
		if c.ID != req.ID && c.ParentID != 0 {
			if parent, exists := categoryMap[c.ParentID]; exists {
				parent.Children = append(parent.Children, c)
			}
		}
	}

	rootVO, err := buildCategoryVOTree(cat, c)
	if err != nil {
		utils.BizLogger(c).Errorf("删除类目时映射 VO 失败: %v", err)
		return nil, fmt.Errorf("删除类目时映射 VO 失败: %w", err)
	}

	return []*category.CategoriesVO{rootVO}, nil
}

// recursivelyUpdateChildrenPaths 递归更新子类目路径
func recursivelyUpdateChildrenPaths(parentCategory *model.Category, c echo.Context) error {
	children, err := mapper.GetCategoriesByParentID(parentCategory.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败: %v", err)
		return fmt.Errorf("获取类目失败: %w", err)
	}

	for _, child := range children {
		if parentCategory.Path == "" {
			child.Path = fmt.Sprintf("/%d", parentCategory.ID)
		} else {
			child.Path = fmt.Sprintf("%s/%d", parentCategory.Path, parentCategory.ID)
		}

		if err := mapper.UpdateCategory(child); err != nil {
			utils.BizLogger(c).Errorf("更新子类目失败: %v", err)
			return fmt.Errorf("更新子类目失败: %w", err)
		}

		if err := recursivelyUpdateChildrenPaths(child, c); err != nil {
			return err
		}
	}

	return nil
}

// buildCategoryVOTree 构建类目树VO
func buildCategoryVOTree(cat *model.Category, c echo.Context) (*category.CategoriesVO, error) {
	vo, err := utils.MapModelToVO(cat, &category.CategoriesVO{})
	if err != nil {
		utils.BizLogger(c).Errorf("映射类目 VO 失败: %v", err)
		return nil, fmt.Errorf("映射类目 VO 失败: %w", err)
	}

	catVO := vo.(*category.CategoriesVO)

	if len(cat.Children) > 0 {
		catVO.Children = make([]*category.CategoriesVO, 0, len(cat.Children))
		for _, child := range cat.Children {
			childVO, err := buildCategoryVOTree(child, c)
			if err != nil {
				return nil, err
			}
			catVO.Children = append(catVO.Children, childVO)
		}
	}

	return catVO, nil
}
