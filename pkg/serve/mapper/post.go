// Package mapper 提供数据模型与数据库交互的映射层，处理文章相关数据操作
// 创建者：Done-0
// 创建时间：2025-05-10
package mapper

import (
	"fmt"

	"github.com/labstack/echo/v4"

	post "jank.com/jank_blog/internal/model/post"
	"jank.com/jank_blog/internal/utils"
)

// CreatePost 将文章保存到数据库
// 参数：
//   - c: Echo 上下文
//   - newPost: 文章信息
//
// 返回值：
//   - error: 操作过程中的错误
func CreatePost(c echo.Context, newPost *post.Post) error {
	db := utils.GetDBFromContext(c)
	if err := db.Create(newPost).Error; err != nil {
		return fmt.Errorf("创建文章失败: %w", err)
	}
	return nil
}

// GetPostByID 根据 ID 获取文章
// 参数：
//   - c: Echo 上下文
//   - id: 文章 ID
//
// 返回值：
//   - *post.Post: 文章信息
//   - error: 操作过程中的错误
func GetPostByID(c echo.Context, id int64) (*post.Post, error) {
	var pos post.Post
	db := utils.GetDBFromContext(c)
	if err := db.Where("id = ? AND deleted = ?", id, false).First(&pos).Error; err != nil {
		return nil, fmt.Errorf("获取文章失败: %w", err)
	}
	return &pos, nil
}

// GetAllPostsWithPaging 获取分页后的文章列表和文章总数
// 参数：
//   - c: Echo 上下文
//   - page: 页码
//   - pageSize: 每页大小
//
// 返回值：
//   - []*post.Post: 文章列表
//   - int64: 文章总数
//   - error: 操作过程中的错误
func GetAllPostsWithPaging(c echo.Context, page, pageSize int) ([]*post.Post, int64, error) {
	var posts []*post.Post
	var total int64
	db := utils.GetDBFromContext(c)

	// 查询文章总数
	if err := db.Model(&post.Post{}).Where("deleted = ?", false).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("获取文章总数失败: %w", err)
	}

	// 使用雪花算法ID排序的分页查询 (雪花ID本身包含时间信息，降序排列即为最新内容)
	if err := db.Where("deleted = ?", false).
		Order("id DESC").
		Limit(pageSize).Offset((page - 1) * pageSize).
		Find(&posts).Error; err != nil {
		return nil, 0, fmt.Errorf("获取分页文章列表失败: %w", err)
	}
	return posts, total, nil
}

// UpdateOnePostByID 更新文章
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//   - newPost: 文章信息
//
// 返回值：
//   - error: 操作过程中的错误
func UpdateOnePostByID(c echo.Context, postID int64, newPost *post.Post) error {
	db := utils.GetDBFromContext(c)
	result := db.Model(&post.Post{}).Where("id = ? AND deleted = ?", postID, false).Updates(newPost)

	if result.Error != nil {
		return fmt.Errorf("更新文章失败: %w", result.Error)
	}
	return nil
}

// DeleteOnePostByID 根据 ID 进行软删除操作
// 参数：
//   - c: Echo 上下文
//   - postID: 文章 ID
//
// 返回值：
//   - error: 操作过程中的错误
func DeleteOnePostByID(c echo.Context, postID int64) error {
	db := utils.GetDBFromContext(c)
	result := db.Model(&post.Post{}).
		Where("id = ? AND deleted = ?", postID, false).
		Update("deleted", true)

	if result.Error != nil {
		return fmt.Errorf("删除文章失败: %w", result.Error)
	}
	return nil
}
