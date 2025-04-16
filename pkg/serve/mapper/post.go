package mapper

import (
	"fmt"

	"jank.com/jank_blog/internal/global"
	post "jank.com/jank_blog/internal/model/post"
)

// CreatePost 将文章保存到数据库
func CreatePost(newPost *post.Post) error {
	if err := global.DB.Create(newPost).Error; err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	}
	return nil
}

// GetPostByID 根据 ID 获取文章
func GetPostByID(id int64) (*post.Post, error) {
	var pos post.Post
	if err := global.DB.Where("id = ? AND deleted = ?", id, false).First(&pos).Error; err != nil {
		return nil, fmt.Errorf("获取文章失败: %w", err)
	}
	return &pos, nil
}

// GetPostsByTitle 通过 Title 获取所有匹配的文章
func GetPostsByTitle(title string) ([]post.Post, error) {
	var posts []post.Post
	if err := global.DB.Where("title LIKE ? AND deleted = ?", "%"+title+"%", false).
		Find(&posts).Error; err != nil {
		return nil, fmt.Errorf("根据标题查询文章失败: %w", err)
	}
	return posts, nil
}

// GetAllPostsWithPaging 获取分页后的文章列表和文章总数
func GetAllPostsWithPaging(page, pageSize int) ([]*post.Post, int64, error) {
	var posts []*post.Post
	var total int64

	offset := (page - 1) * pageSize

	// 查询文章总数
	if err := global.DB.Model(&post.Post{}).Where("deleted = ?", false).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("获取文章总数失败: %w", err)
	}

	// 查询分页数据
	if err := global.DB.Where("deleted = ?", false).
		Order("gmt_create DESC").
		Offset(offset).Limit(pageSize).
		Find(&posts).Error; err != nil {
		return nil, 0, fmt.Errorf("获取分页文章列表失败: %w", err)
	}
	return posts, total, nil
}

// UpdateOnePostByID 更新文章
func UpdateOnePostByID(postID int64, newPost *post.Post) error {
	result := global.DB.Model(&post.Post{}).Where("id = ? AND deleted = ?", postID, false).Updates(newPost)

	if result.Error != nil {
		return fmt.Errorf("更新文章失败: %w", result.Error)
	}
	return nil
}

// DeleteOnePostByID 根据 ID 进行软删除操作
func DeleteOnePostByID(postID int64) error {
	result := global.DB.Model(&post.Post{}).
		Where("id = ? AND deleted = ?", postID, false).
		Update("deleted", true)

	if result.Error != nil {
		return fmt.Errorf("删除文章失败: %w", result.Error)
	}
	return nil
}
