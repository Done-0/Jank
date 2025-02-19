package service

import (
	"fmt"
	"jank.com/jank_blog/pkg/serve/controller/category/dto"

	"github.com/labstack/echo/v4"

	model "jank.com/jank_blog/internal/model/category"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/mapper"
	"jank.com/jank_blog/pkg/vo/category"
)

// GetCategoryByID 根据 ID 获取类目
func GetCategoryByID(req *dto.GetOneCategoryRequest, c echo.Context) (*category.CategoriesVo, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%v", err)
	}

	vo, err := utils.MapModelToVO(cat, &category.CategoriesVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("获取类目时映射 vo 失败：%v", err)
	}

	return vo.(*category.CategoriesVo), nil
}

// GetCategoryTree 获取类目树
func GetCategoryTree(c echo.Context) ([]category.CategoriesVo, error) {
	categories, err := mapper.GetAllActivatedCategories()
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目树失败: %v", err)
	}

	categoryMap := make(map[int64]*model.Category)
	var rootCategories []*model.Category

	// 构建类目映射（ID -> 类目对象指针）
	for i := range categories {
		categoryMap[categories[i].ID] = categories[i]
	}

	// 遍历类目，构建父子关系
	for i := range categories {
		cat := categories[i]

		if cat.ParentID == 0 {
			rootCategories = append(rootCategories, cat)
		} else {
			parentCategory, exists := categoryMap[cat.ParentID]
			if exists {
				if parentCategory.Children == nil {
					parentCategory.Children = make([]*model.Category, 0)
				}
				parentCategory.Children = append(parentCategory.Children, cat)
			}
		}
	}

	var categoryTree []category.CategoriesVo
	for _, root := range rootCategories {
		categoryVo, err := utils.MapModelToVO(root, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取类目树时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("获取类目树时映射 vo 失败：%v", err)
		}
		categoryTree = append(categoryTree, *categoryVo.(*category.CategoriesVo))
	}

	return categoryTree, nil
}

// GetCategoryChildrenByID 根据类目 ID 获取层级子类目
func GetCategoryChildrenByID(req *dto.GetOneCategoryRequest, c echo.Context) ([]category.CategoriesVo, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%v", err)
	}

	// 获取该类目下所有子类目
	children, err := recursivelyGetChildren(cat.ID, c)
	if err != nil {
		utils.BizLogger(c).Errorf("获取层级子类目失败：%v", err)
		return nil, fmt.Errorf("获取层级子类目失败：%v", err)
	}

	// 将子类目映射为 CategoriesVo
	var childrenVo []category.CategoriesVo
	for _, child := range children {
		childVo, err := utils.MapModelToVO(child, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取子类目时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("获取子类目时映射 vo 失败：%v", err)
		}
		childrenVo = append(childrenVo, *childVo.(*category.CategoriesVo))
	}

	return childrenVo, nil
}

// CreateCategory 创建类目
func CreateCategory(req *dto.CreateOneCategoryRequest, c echo.Context) (*category.CategoriesVo, error) {
	var newCategory *model.Category

	// 创建根类目
	if req.ParentID == 0 {
		newCategory = &model.Category{
			Name:        req.Name,
			Description: req.Description,
			ParentID:    0,
			Path:        "",
		}

		// 创建根类目
		if err := mapper.CreateCategory(newCategory); err != nil {
			utils.BizLogger(c).Errorf("创建根类目失败：%v", err)
			return nil, fmt.Errorf("创建根类目失败: %v", err)
		}

		// 映射为 VO 类型
		categoryVo, err := utils.MapModelToVO(newCategory, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("创建类目时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("创建类目时映射 vo 失败：%v", err)
		}
		return categoryVo.(*category.CategoriesVo), nil
	}

	// 获取父类目
	cat, err := mapper.GetCategoryByID(req.ParentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取父类目失败：%v", err)
		return nil, fmt.Errorf("获取父类目失败：%v", err)
	}

	// 创建子类目
	newCategory = &model.Category{
		Name:        req.Name,
		Description: req.Description,
		ParentID:    req.ParentID,
		Path:        fmt.Sprintf("%s/%d", cat.Path, req.ParentID),
	}

	if err := mapper.CreateCategory(newCategory); err != nil {
		utils.BizLogger(c).Errorf("创建子类目失败：%v", err)
		return nil, fmt.Errorf("创建子类目失败: %v", err)
	}

	categoryVo, err := utils.MapModelToVO(newCategory, &category.CategoriesVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("创建类目时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("创建类目时映射 vo 失败：%v", err)
	}
	return categoryVo.(*category.CategoriesVo), nil
}

// UpdateCategory 更新类目
func UpdateCategory(req *dto.UpdateOneCategoryRequest, c echo.Context) (*category.CategoriesVo, error) {
	existingCategory, err := mapper.GetCategoryByID(req.ParentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败: %v", err)
	}

	parentPath, err := mapper.GetParentCategoryPathByID(req.ParentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取[%s]父类目路径失败：%v", existingCategory.Name, err)
		return nil, fmt.Errorf("获取[%s]父类目路径失败：%v", existingCategory.Name, err)
	}

	existingCategory.Name = req.Name
	existingCategory.Description = req.Description
	existingCategory.ParentID = req.ParentID
	existingCategory.Path = fmt.Sprintf("%s/%d", parentPath, req.ParentID)

	if err := mapper.UpdateCategory(existingCategory); err != nil {
		utils.BizLogger(c).Errorf("[%s]类目更新失败：%v", existingCategory.Name, err)
		return nil, fmt.Errorf("[%s]类目更新失败: %v", existingCategory.Name, err)
	}

	if err := recursivelyUpdateChildrenPaths(existingCategory, c); err != nil {
		utils.BizLogger(c).Errorf("递归更新[%s]类目失败: %v", existingCategory.Name, err)
		return nil, fmt.Errorf("递归更新[%s]类目失败: %v", existingCategory.Name, err)
	}

	categoryVo, err := utils.MapModelToVO(existingCategory, &category.CategoriesVo{})
	if err != nil {
		utils.BizLogger(c).Errorf("更新类目时映射 vo 失败：%v", err)
		return nil, fmt.Errorf("更新类目时映射 vo 失败：%v", err)
	}

	return categoryVo.(*category.CategoriesVo), nil
}

// DeleteCategory 软删除类目
func DeleteCategory(req *dto.DeleteOneCategoryRequest, c echo.Context) ([]category.CategoriesVo, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%v", err)
	}

	deletedCategories, err := mapper.GetCategoriesByPath(cat.Path)
	if err != nil {
		utils.BizLogger(c).Errorf("获取[%s]下所有子类目失败：%v", cat.Name, err)
		return nil, fmt.Errorf("获取[%s]下所有子类目失败：%v", cat.Name, err)
	}

	if err := mapper.DeleteCategoriesByPathSoftly(cat.Path); err != nil {
		utils.BizLogger(c).Errorf("软删除[%s]下所有子类目失败：%v", cat.Name, err)
		return nil, fmt.Errorf("软删除[%s]下所有子类目失败：%v", cat.Name, err)
	}

	var deletedCategoriesVo []category.CategoriesVo
	for _, deletedCategory := range deletedCategories {
		categoryVo, err := utils.MapModelToVO(deletedCategory, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("删除类目时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("删除类目时映射 vo 失败：%v", err)
		}
		deletedCategoriesVo = append(deletedCategoriesVo, *categoryVo.(*category.CategoriesVo))
	}

	return deletedCategoriesVo, nil
}

// recursivelyUpdateChildrenPaths 递归更新子类目路径
func recursivelyUpdateChildrenPaths(parentCategory *model.Category, c echo.Context) error {
	children, err := mapper.GetCategoriesByParentID(parentCategory.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("递归时获取[%s]父类目失败：%v", parentCategory.Name, err)
		return nil
	}

	// 更新每个子类目的路径
	for _, child := range children {
		child.Path = fmt.Sprintf("%s/%d", parentCategory.Path, child.ID)

		if err := mapper.UpdateCategory(&child); err != nil {
			return fmt.Errorf("更新[%s]子类目失败：%v", child.Name, err)
		}

		if err := recursivelyUpdateChildrenPaths(&child, c); err != nil {
			return fmt.Errorf("递归更新[%s]子类目失败：%v", child.Name, err)
		}
	}

	return nil
}

// recursivelyGetChildren 递归获取子类目
func recursivelyGetChildren(parentID int64, c echo.Context) ([]*model.Category, error) {
	categories, err := mapper.GetCategoriesByParentID(parentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取子类目失败：%v", err)
		return nil, fmt.Errorf("获取子类目失败：%v", err)
	}

	var result []*model.Category

	// 遍历每个子类目，递归获取其子类目
	for _, cat := range categories {
		children, err := recursivelyGetChildren(cat.ID, c)
		if err != nil {
			utils.BizLogger(c).Errorf("递归获取[%s]子类目失败：%v", cat.Name, err)
			return nil, fmt.Errorf("递归获取[%s]子类目失败：%v", cat.Name, err)
		}

		cat.Children = children
		result = append(result, &cat)
	}

	return result, nil
}
