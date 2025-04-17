package mapper

import (
	"errors"
	"fmt"

	"gorm.io/gorm"

	"jank.com/jank_blog/internal/global"
	association "jank.com/jank_blog/internal/model/association"
)

// CreatePostCategory 创建文章-类目关联
func CreatePostCategory(postID, categoryID int64) error {
	postCategory := &association.PostCategory{
		PostID:     postID,
		CategoryID: categoryID,
	}
	if err := global.DB.Create(postCategory).Error; err != nil {
		return fmt.Errorf("创建文章-类目关联失败: %w", err)
	}
	return nil
}

// GetPostCategory 获取文章-类目关联
func GetPostCategory(postID int64) (*association.PostCategory, error) {
	var postCategory association.PostCategory
	err := global.DB.Where("post_id = ? AND deleted = ?", postID, false).First(&postCategory).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("获取文章-类目关联失败: %w", err)
		}
		return nil, fmt.Errorf("获取文章-类目关联失败: %w", err)
	}
	return &postCategory, nil
}

// UpdatePostCategory 更新文章-类目关联
func UpdatePostCategory(postID, categoryID int64) error {
	var exists int64
	if err := global.DB.Model(&association.PostCategory{}).
		Where("post_id = ? AND deleted = ?", postID, false).
		Count(&exists).Error; err != nil {
		return fmt.Errorf("检查文章-类目关联失败: %w", err)
	}
	if exists > 0 {
		if err := global.DB.Model(&association.PostCategory{}).
			Where("post_id = ? AND deleted = ?", postID, false).
			Update("category_id", categoryID).Error; err != nil {
			return fmt.Errorf("更新文章-类目关联失败: %w", err)
		}
	} else {
		return CreatePostCategory(postID, categoryID)
	}

	return nil
}

// DeletePostCategory 删除文章-类目关联
func DeletePostCategory(postID int64) error {
	if err := global.DB.Model(&association.PostCategory{}).
		Where("post_id = ? AND deleted = ?", postID, false).
		Update("deleted", true).Error; err != nil {
		return fmt.Errorf("删除文章-类目关联失败: %w", err)
	}
	return nil
}

// DeletePostCategoryByCategoryID 根据类目ID删除文章-类目关联
func DeletePostCategoryByCategoryID(categoryID int64) error {
	if err := global.DB.Model(&association.PostCategory{}).
		Where("category_id = ? AND deleted = ?", categoryID, false).
		Update("deleted", true).Error; err != nil {
		return fmt.Errorf("根据类目ID删除文章-类目关联失败: %w", err)
	}
	return nil
}
