package category

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
func GetCategoryByID(req *dto.GetOneCategoryRequest, c echo.Context) (*model.Category, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%w", err)
	}
	return cat, nil
}

// GetCategoryTree 获取类目树
func GetCategoryTree(c echo.Context) ([]*category.CategoriesVo, error) {
	categories, err := mapper.GetAllActivatedCategories()
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目树失败：%v", err)
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
		} else {
			if parent, exists := categoryMap[cat.ParentID]; exists {
				if parent.Children == nil {
					parent.Children = make([]*model.Category, 0)
				}
				parent.Children = append(parent.Children, cat)
			}
		}
	}

	var rootCategoriesVo []*category.CategoriesVo
	for _, root := range rootCategories {
		rootCategoryVo, err := utils.MapModelToVO(root, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("获取类目树时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("获取类目树时映射 vo 失败：%v", err)
		}

		if root.Children != nil {
			rootCategoryVo.(*category.CategoriesVo).Children = make([]*category.CategoriesVo, 0)
			for _, child := range root.Children {
				childVo, err := utils.MapModelToVO(child, &category.CategoriesVo{})
				if err != nil {
					utils.BizLogger(c).Errorf("获取类目树时映射子类目 vo 失败：%v", err)
					return nil, fmt.Errorf("获取类目树时映射子类目 vo 失败：%v", err)
				}
				if child.Children != nil {
					childVo.(*category.CategoriesVo).Children = make([]*category.CategoriesVo, 0)
					for _, subChild := range child.Children {
						subChildVo, err := utils.MapModelToVO(subChild, &category.CategoriesVo{})
						if err != nil {
							utils.BizLogger(c).Errorf("获取类目树时映射孙类目 vo 失败：%v", err)
							return nil, fmt.Errorf("获取类目树时映射孙类目 vo 失败：%v", err)
						}
						childVo.(*category.CategoriesVo).Children = append(childVo.(*category.CategoriesVo).Children, subChildVo.(*category.CategoriesVo))
					}
				}
				rootCategoryVo.(*category.CategoriesVo).Children = append(rootCategoryVo.(*category.CategoriesVo).Children, childVo.(*category.CategoriesVo))
			}
		}

		rootCategoriesVo = append(rootCategoriesVo, rootCategoryVo.(*category.CategoriesVo))
	}
	return rootCategoriesVo, nil
}

// GetCategoryChildrenByID 根据类目 ID 获取层级子类目
func GetCategoryChildrenByID(req *dto.GetOneCategoryRequest, c echo.Context) ([]*category.CategoriesVo, error) {
	categories, err := mapper.GetAllActivatedCategories()
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目树失败：%v", err)
		return nil, fmt.Errorf("获取类目树失败: %w", err)
	}

	// 构建类目映射
	categoryMap := make(map[int64]*model.Category)
	for i := range categories {
		categoryMap[categories[i].ID] = categories[i]
	}

	// 构建父子关系
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

	// 找到目标类目
	target, exists := categoryMap[req.ID]
	if !exists {
		return nil, fmt.Errorf("未找到类目 ID: %d", req.ID)
	}

	var childrenVo []*category.CategoriesVo
	if target.Children != nil {
		for _, child := range target.Children {
			childVo, err := utils.MapModelToVO(child, &category.CategoriesVo{})
			if err != nil {
				utils.BizLogger(c).Errorf("获取层级子类目时映射 vo 失败：%v", err)
				return nil, fmt.Errorf("获取层级子类目时映射 vo 失败：%v", err)
			}

			if child.Children != nil {
				childVo.(*category.CategoriesVo).Children = make([]*category.CategoriesVo, 0)
				for _, subChild := range child.Children {
					subChildVo, err := utils.MapModelToVO(subChild, &category.CategoriesVo{})
					if err != nil {
						utils.BizLogger(c).Errorf("获取层级子类目时映射孙类目 vo 失败：%v", err)
						return nil, fmt.Errorf("获取层级子类目时映射孙类目 vo 失败：%v", err)
					}
					childVo.(*category.CategoriesVo).Children = append(childVo.(*category.CategoriesVo).Children, subChildVo.(*category.CategoriesVo))
				}
			}
			childrenVo = append(childrenVo, childVo.(*category.CategoriesVo))
		}
	}
	return childrenVo, nil
}

// CreateCategory 创建类目
func CreateCategory(req *dto.CreateOneCategoryRequest, c echo.Context) (*category.CategoriesVo, error) {
	var newCategory *model.Category

	if req.ParentID == 0 {
		newCategory = &model.Category{
			Name:        req.Name,
			Description: req.Description,
			ParentID:    0,
			Path:        "",
		}

		if err := mapper.CreateCategory(newCategory); err != nil {
			utils.BizLogger(c).Errorf("创建根类目失败：%v", err)
			return nil, fmt.Errorf("创建根类目失败: %v", err)
		}

		categoryVo, err := utils.MapModelToVO(newCategory, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("创建类目时映射 vo 失败：%v", err)
			return nil, fmt.Errorf("创建类目时映射 vo 失败：%v", err)
		}
		return categoryVo.(*category.CategoriesVo), nil
	}

	cat, err := mapper.GetCategoryByID(req.ParentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取父类目失败：%v", err)
		return nil, fmt.Errorf("获取父类目失败：%v", err)
	}

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
	existingCategory, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败: %v", err)
	}

	parentPath, err := mapper.GetParentCategoryPathByID(req.ParentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取「%v」父类目路径失败：%v", existingCategory.Name, err)
		return nil, fmt.Errorf("获取「%v」父类目路径失败：%v", existingCategory.Name, err)
	}

	existingCategory.Name = req.Name
	existingCategory.Description = req.Description
	existingCategory.ParentID = req.ParentID
	existingCategory.Path = fmt.Sprintf("%s/%d", parentPath, req.ParentID)

	if err := mapper.UpdateCategory(existingCategory); err != nil {
		utils.BizLogger(c).Errorf("「%v」类目更新失败：%v", existingCategory.Name, err)
		return nil, fmt.Errorf("「%v」类目更新失败: %v", existingCategory.Name, err)
	}

	if err := recursivelyUpdateChildrenPaths(existingCategory, c); err != nil {
		utils.BizLogger(c).Errorf("递归更新「%v」类目失败: %v", existingCategory.Name, err)
		return nil, fmt.Errorf("递归更新「%v」类目失败: %v", existingCategory.Name, err)
	}

	var convert func(cat *model.Category) (*category.CategoriesVo, error)
	convert = func(cat *model.Category) (*category.CategoriesVo, error) {
		vo, err := utils.MapModelToVO(cat, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("映射类目 VO 失败：%v", err)
			return nil, fmt.Errorf("映射类目 VO 失败：%v", err)
		}
		catVo := vo.(*category.CategoriesVo)
		if cat.Children != nil && len(cat.Children) > 0 {
			catVo.Children = make([]*category.CategoriesVo, 0, len(cat.Children))
			for _, child := range cat.Children {
				childVo, err := convert(child)
				if err != nil {
					return nil, err
				}
				catVo.Children = append(catVo.Children, childVo)
			}
		}
		return catVo, nil
	}

	updatedVo, err := convert(existingCategory)
	if err != nil {
		return nil, err
	}
	return updatedVo, nil
}

// DeleteCategory 软删除类目
func DeleteCategory(req *dto.DeleteOneCategoryRequest, c echo.Context) ([]*category.CategoriesVo, error) {
	cat, err := mapper.GetCategoryByID(req.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%w", err)
	}

	deletedCategories, err := mapper.GetCategoriesByPath(cat.Path)
	if err != nil {
		utils.BizLogger(c).Errorf("获取「%v」下所有子类目失败：%v", cat.Name, err)
		return nil, fmt.Errorf("获取「%v」下所有子类目失败：%v", cat.Name, err)
	}

	if err := mapper.DeleteCategoriesByPathSoftly(cat.Path, req.ID); err != nil {
		utils.BizLogger(c).Errorf("软删除「%v」下所有子类目失败：%v", cat.Name, err)
		return nil, fmt.Errorf("软删除「%v」下所有子类目失败：%v", cat.Name, err)
	}

	var convert func(cat *model.Category) (*category.CategoriesVo, error)
	convert = func(cat *model.Category) (*category.CategoriesVo, error) {
		vo, err := utils.MapModelToVO(cat, &category.CategoriesVo{})
		if err != nil {
			utils.BizLogger(c).Errorf("映射类目 Vo 失败：%v", err)
			return nil, fmt.Errorf("映射类目 Vo 失败：%v", err)
		}
		catVo := vo.(*category.CategoriesVo)

		if cat.Children != nil && len(cat.Children) > 0 {
			catVo.Children = make([]*category.CategoriesVo, 0, len(cat.Children))
			for _, child := range cat.Children {
				childVo, err := convert(child)
				if err != nil {
					return nil, err
				}
				catVo.Children = append(catVo.Children, childVo)
			}
		}
		return catVo, nil
	}

	var deletedCategoriesVo []*category.CategoriesVo
	for _, deletedCategory := range deletedCategories {
		deletedCategoryVo, err := convert(deletedCategory)
		if err != nil {
			utils.BizLogger(c).Errorf("映射删除的类目时失败：%v", err)
			return nil, fmt.Errorf("映射删除的类目时失败：%v", err)
		}
		deletedCategoriesVo = append(deletedCategoriesVo, deletedCategoryVo)
	}

	return deletedCategoriesVo, nil
}

// recursivelyUpdateChildrenPaths 递归更新子类目路径
func recursivelyUpdateChildrenPaths(parentCategory *model.Category, c echo.Context) error {
	children, err := mapper.GetCategoriesByParentID(parentCategory.ID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取「%v」子类目失败：%v", parentCategory.Name, err)
		utils.BizLogger(c).Errorf("递归时获取「%v」父类目失败：%v", parentCategory.Name, err)
		return nil
	}

	// 更新每个子类目的路径
	for _, child := range children {
		child.Path = fmt.Sprintf("%s/%d", parentCategory.Path, child.ID)

		if err := mapper.UpdateCategory(child); err != nil {
			utils.BizLogger(c).Errorf("更新「%v」子类目失败：%v", child.Name, err)
			return fmt.Errorf("更新「%v」子类目失败：%v", child.Name, err)
		}

		if err := recursivelyUpdateChildrenPaths(child, c); err != nil {
			utils.BizLogger(c).Errorf("递归更新「%v」子类目失败：%v", child.Name, err)
			return fmt.Errorf("递归更新「%v」子类目失败：%v", child.Name, err)
		}
	}

	return nil
}
