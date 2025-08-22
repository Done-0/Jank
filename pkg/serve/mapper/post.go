// Package mapper 提供文章相关的数据访问接口
// 创建者：Done-0
// 创建时间：2025-08-13
package mapper

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/post"
)

// PostMapper 文章数据访问接口
type PostMapper interface {
	GetPostByID(c *app.RequestContext, postID int64) (*post.Post, error)                                                            // 根据 ID 获取文章
	ListPublishedPosts(c *app.RequestContext, pageNo, pageSize int64, categoryID *int64) ([]*post.Post, int64, error)               // 获取已发布文章列表，categoryID为空时不按分类筛选
	ListPostsByStatus(c *app.RequestContext, pageNo, pageSize int64, status string, categoryID *int64) ([]*post.Post, int64, error) // 根据状态获取文章列表，status为空时获取所有文章，categoryID为空时不按分类筛选
	ListPublicPosts(c *app.RequestContext, pageNo, pageSize int64) ([]*post.Post, int64, error)                                     // 获取公开文章（已发布+已归档）
	CreatePost(c *app.RequestContext, post *post.Post) error                                                                        // 创建文章
	UpdatePost(c *app.RequestContext, post *post.Post) error                                                                        // 更新文章
	DeletePost(c *app.RequestContext, postID int64) error                                                                           // 删除文章
}
