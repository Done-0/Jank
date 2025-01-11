package mapper

import (
	"fmt"
	"strings"

	"jank.com/jank_blog/internal/global"
	category "jank.com/jank_blog/internal/model/category"
	post "jank.com/jank_blog/internal/model/post"
)

// getValidCategoryIDs 获取未删除的分类 ID 列表并更新数据库
func getValidCategoryIDs(postID int64, categoryIDs string) (string, bool, error) {
	if categoryIDs == "" {
		return "", false, nil
	}

	ids := strings.Split(categoryIDs, ",")
	var validIDs []string
	updated := false

	for _, id := range ids {
		var category category.Category
		err := global.DB.Where("id = ? AND deleted = ?", id, 0).First(&category).Error
		if err == nil {
			validIDs = append(validIDs, id)
		} else {
			updated = true
		}
	}

	newCategoryIDs := strings.Join(validIDs, ",")

	if updated && postID > 0 {
		err := global.DB.Model(&post.Post{}).Where("id = ?", postID).Update("category_ids", newCategoryIDs).Error
		if err != nil {
			return "", false, err
		}
	}

	return newCategoryIDs, updated, nil
}

// CreatePost 将文章保存到数据库
func CreatePost(newPost *post.Post) error {
	if newPost == nil {
		return fmt.Errorf("文章不能为空")
	}

	validCategoryIDs, _, err := getValidCategoryIDs(0, newPost.CategoryIDs)
	if err != nil {
		return err
	}
	newPost.CategoryIDs = validCategoryIDs

	return global.DB.Create(newPost).Error
}

// GetPostByID 根据 ID 获取文章
func GetPostByID(id int64) (*post.Post, error) {
	if id <= 0 {
		return nil, fmt.Errorf("无效文章ID: %d", id)
	}

	var post post.Post
	err := global.DB.Where("id = ? AND deleted = ?", id, 0).First(&post).Error
	if err != nil {
		return nil, err
	}

	validCategoryIDs, updated, err := getValidCategoryIDs(post.ID, post.CategoryIDs)
	if err != nil {
		return nil, err
	}
	post.CategoryIDs = validCategoryIDs

	if updated {
		err = global.DB.Save(&post).Error
		if err != nil {
			return nil, err
		}
	}

	return &post, nil
}

// GetPostsByTitle 通过 Title 获取所有匹配的文章
func GetPostsByTitle(title string) ([]post.Post, error) {
	if title == "" {
		return nil, fmt.Errorf("文章标题不能为空")
	}

	var posts []post.Post
	err := global.DB.Where("title LIKE ? AND deleted = ?", "%"+title+"%", 0).
		Find(&posts).Error

	if err != nil {
		return nil, err
	}

	for i := range posts {
		validCategoryIDs, updated, err := getValidCategoryIDs(posts[i].ID, posts[i].CategoryIDs)
		if err != nil {
			return nil, err
		}
		posts[i].CategoryIDs = validCategoryIDs
		if updated {
			err = global.DB.Save(&posts[i]).Error
			if err != nil {
				return nil, err
			}
		}
	}

	return posts, nil
}

// GetAllPostsWithPaging 获取分页后的文章列表和文章总数
func GetAllPostsWithPaging(page, pageSize int) ([]*post.Post, int64, error) {
	var posts []*post.Post
	var total int64

	offset := (page - 1) * pageSize

	// 查询文章总数
	err := global.DB.Model(&post.Post{}).Where("deleted = ?", 0).Count(&total).Error
	if err != nil {
		return nil, 0, err
	}

	// 查询分页数据
	err = global.DB.Where("deleted = ?", 0).
		Order("gmt_create DESC").
		Offset(offset).Limit(pageSize).
		Find(&posts).Error
	if err != nil {
		return nil, 0, err
	}

	// 文章类别 ID 列表更新
	for i, post := range posts {
		validCategoryIDs, updated, err := getValidCategoryIDs(post.ID, post.CategoryIDs)
		if err != nil {
			return nil, 0, err
		}
		posts[i].CategoryIDs = validCategoryIDs
		if updated {
			err = global.DB.Save(posts[i]).Error
			if err != nil {
				return nil, 0, err
			}
		}
	}

	return posts, total, nil
}

// UpdateOnePostByID 更新文章
func UpdateOnePostByID(postID int64, newPost *post.Post) error {
	if postID <= 0 || newPost == nil {
		return fmt.Errorf("无效文章ID: %d", postID)
	}

	validCategoryIDs, _, err := getValidCategoryIDs(postID, newPost.CategoryIDs)
	if err != nil {
		return err
	}
	newPost.CategoryIDs = validCategoryIDs

	result := global.DB.Model(&post.Post{}).Where("id = ? AND deleted = ?", postID, 0).Updates(newPost)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在或已经删除")
	}

	return nil
}

// DeleteOnePostByID 根据 ID 进行软删除操作
func DeleteOnePostByID(postID int64) error {
	if postID <= 0 {
		return fmt.Errorf("无效文章ID: %d", postID)
	}

	result := global.DB.Model(&post.Post{}).
		Where("id = ? AND deleted = ?", postID, 0).
		Update("deleted", true)

	if result.Error != nil {
		return result.Error
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("文章不存在或已经删除")
	}

	return nil
}
