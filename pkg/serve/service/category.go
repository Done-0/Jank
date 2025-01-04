package service

import (
	"fmt"

	"github.com/labstack/echo/v4"
	model "jank.com/jank_blog/internal/model/category"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/serve/mapper"
)

// GetCategoryByID 根据 ID 获取类目
func GetCategoryByID(id int64, c echo.Context) (*model.Category, error) {
	category, err := mapper.GetCategoryByID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%w", err)
	}
	return category, nil
}

// GetCategoryTree 获取类目树
func GetCategoryTree() ([]model.Category, error) {
	categories, err := mapper.GetAllActivedCategories()
	if err != nil {
		return nil, fmt.Errorf("获取类目树失败: %w", err)
	}

	categoryMap := make(map[int64]*model.Category)
	var rootCategories []*model.Category

	// 构建类目映射（ID -> 类目对象指针）
	for i := range categories {
		categoryMap[categories[i].ID] = categories[i]
	}

	// 遍历类目，构建父子关系
	for i := range categories {
		category := categories[i]

		if category.ParentID == 0 {
			rootCategories = append(rootCategories, category)
		} else {
			parentCategory, exists := categoryMap[category.ParentID]
			if exists {
				if parentCategory.Children == nil {
					parentCategory.Children = make([]*model.Category, 0)
				}
				parentCategory.Children = append(parentCategory.Children, category)
			}
		}
	}

	var categoryTree []model.Category
	for _, root := range rootCategories {
		categoryTree = append(categoryTree, *root)
	}

	return categoryTree, nil
}

// GetCategoryChildrenByID 根据类目 ID 获取层级子类目
func GetCategoryChildrenByID(id int64, c echo.Context) ([]*model.Category, error) {
	category, err := mapper.GetCategoryByID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%v", err)
	}

	// 获取该类目下所有子类目
	children, err := recursivelyGetChildren(category.ID, c)
	if err != nil {
		utils.BizLogger(c).Errorf("获取层级子类目失败：%v", err)
		return nil, fmt.Errorf("获取层级子类目失败：%w", err)
	}

	return children, nil
}

// CreateCategory 创建类目
func CreateCategory(name string, description string, parentID int64, c echo.Context) (*model.Category, error) {
    var newCategory *model.Category

    // 创建根类目
    if parentID == 0 {
        newCategory = &model.Category{
            Name:        name,
            Description: description,
            ParentID:    0,
            Path:        "",
        }

        // 创建根类目
        if err := mapper.CreateCategory(newCategory); err != nil {
            utils.BizLogger(c).Errorf("创建根类目失败：%v", err)
            return nil, fmt.Errorf("创建根类目失败: %w", err)
        }

        return newCategory, nil
    }

    // 获取父类目
    category, err := mapper.GetCategoryByID(parentID)
    if err != nil {
        utils.BizLogger(c).Errorf("获取父类目失败：%v", err)
        return nil, fmt.Errorf("获取父类目失败: %w", err)
    }

    // 创建子类目
    newCategory = &model.Category{
        Name:        name,
        Description: description,
        ParentID:    parentID,
        Path:        fmt.Sprintf("%s/%d", category.Path, parentID), 
    }

    if err := mapper.CreateCategory(newCategory); err != nil {
        utils.BizLogger(c).Errorf("创建子类目失败：%v", err)
        return nil, fmt.Errorf("创建子类目失败: %w", err)
    }

    return newCategory, nil
}

// UpdateCategory 更新类目
func UpdateCategory(id int64, name string, description string, parentID int64, c echo.Context) (*model.Category, error) {
	existingCategory, err := mapper.GetCategoryByID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败: %w", err)
	}

	parentPath, err := mapper.GetParentCategoryPathByID(parentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取[%s]父类目路径失败：%v", existingCategory.Name, err)
		return nil, fmt.Errorf("获取[%s]父类目路径失败：%v", existingCategory.Name, err)
	}

	existingCategory.Name = name
	existingCategory.Description = description
	existingCategory.ParentID = parentID
	existingCategory.Path = fmt.Sprintf("%s/%d", parentPath, parentID)

	if err := mapper.UpdateCategory(existingCategory); err != nil {
		return nil, fmt.Errorf("[%s]类目更新失败: %w", existingCategory.Name, err)
	}

	if err := recursivelyUpdateChildrenPaths(existingCategory, c); err != nil {
		return nil, err
	}

	return existingCategory, nil
}

// DeleteCategory 软删除类目
func DeleteCategory(id int64, c echo.Context) ([]*model.Category, error) {
	category, err := mapper.GetCategoryByID(id)
	if err != nil {
		utils.BizLogger(c).Errorf("获取类目失败：%v", err)
		return nil, fmt.Errorf("获取类目失败：%w", err)
	}

	deletedCategories, err := mapper.GetCategoriesByPath(category.Path)
	if err != nil {
		utils.BizLogger(c).Errorf("获取[%s]下所有子类目失败失败：%v", category.Name, err)
		return nil, fmt.Errorf("获取[%s]下所有子类目失败：%v", category.Name, err)
	}

	if err := mapper.DeleteCategoriesByPathSoftly(category.Path); err != nil {
		utils.BizLogger(c).Errorf("软删除[%s]下所有子类目失败：%v", category.Name, err)
		return nil, fmt.Errorf("软删除[%s]下所有子类目失败：%v", category.Name, err)
	}

	return deletedCategories, nil
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
			return err
		}

		if err := recursivelyUpdateChildrenPaths(&child, c); err != nil {
			return err
		}
	}

	return nil
}

// recursivelyGetChildren 递归获取子类目
func recursivelyGetChildren(parentID int64, c echo.Context) ([]*model.Category, error) {
	categories, err := mapper.GetCategoriesByParentID(parentID)
	if err != nil {
		utils.BizLogger(c).Errorf("获取子类目失败：%v", err)
		return nil, fmt.Errorf("获取子类目失败：%w", err)
	}

	var result []*model.Category

	// 遍历每个子类目，递归获取其子类目
	for _, category := range categories {
		children, err := recursivelyGetChildren(category.ID, c)
		if err != nil {
			utils.BizLogger(c).Errorf("递归获取[%s]子类目失败：%v", category.Name, err)
			return nil, fmt.Errorf("递归获取[%s]子类目失败：%v", category.Name, err)
		}

		category.Children = children
		result = append(result, &category)
	}

	return result, nil
}
