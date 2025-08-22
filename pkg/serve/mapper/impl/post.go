// Package impl 提供文章相关的数据访问实现
// 创建者：Done-0
// 创建时间：2025-08-13
package impl

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/post"
	"github.com/Done-0/jank/internal/types/consts"
	"github.com/Done-0/jank/internal/utils/db"
	"github.com/Done-0/jank/pkg/serve/mapper"
)

// PostMapperImpl 文章数据访问实现
type PostMapperImpl struct{}

// NewPostMapper 创建文章数据访问实例
func NewPostMapper() mapper.PostMapper {
	return &PostMapperImpl{}
}

// GetPostByID 根据ID获取文章
func (m *PostMapperImpl) GetPostByID(c *app.RequestContext, postID int64) (*post.Post, error) {
	var p post.Post
	err := db.GetDBFromContext(c).Where("id = ? AND deleted = ?", postID, false).First(&p).Error
	if err != nil {
		return nil, err
	}
	return &p, nil
}

// ListPublishedPosts 获取已发布文章列表，categoryID 为空时不按分类筛选
func (m *PostMapperImpl) ListPublishedPosts(c *app.RequestContext, pageNo, pageSize int64, categoryID *int64) ([]*post.Post, int64, error) {
	var posts []*post.Post
	var total int64

	query := db.GetDBFromContext(c).Model(&post.Post{}).Where("deleted = ? AND status = ?", false, consts.PostStatusPublished)
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.Order("id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// ListPostsByStatus 根据状态获取文章列表，status 为空时获取所有文章，categoryID 为空时不按分类筛选
func (m *PostMapperImpl) ListPostsByStatus(c *app.RequestContext, pageNo, pageSize int64, status string, categoryID *int64) ([]*post.Post, int64, error) {
	var posts []*post.Post
	var total int64

	query := db.GetDBFromContext(c).Model(&post.Post{}).Where("deleted = ?", false)
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if categoryID != nil {
		query = query.Where("category_id = ?", *categoryID)
	}

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.Order("id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// ListPublicPosts 获取公开文章（已发布+已归档）
func (m *PostMapperImpl) ListPublicPosts(c *app.RequestContext, pageNo, pageSize int64) ([]*post.Post, int64, error) {
	var posts []*post.Post
	var total int64

	// 查询已发布和已归档的文章
	query := db.GetDBFromContext(c).Model(&post.Post{}).Where("deleted = ? AND status IN (?, ?)", false, consts.PostStatusPublished, consts.PostStatusArchived)

	// 统计总数
	if err := query.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	// 分页查询
	offset := (pageNo - 1) * pageSize
	if err := query.Order("id DESC").Offset(int(offset)).Limit(int(pageSize)).Find(&posts).Error; err != nil {
		return nil, 0, err
	}

	return posts, total, nil
}

// CreatePost 创建文章
func (m *PostMapperImpl) CreatePost(c *app.RequestContext, p *post.Post) error {
	if err := db.GetDBFromContext(c).Create(p).Error; err != nil {
		return err
	}
	return nil
}

// UpdatePost 更新文章
func (m *PostMapperImpl) UpdatePost(c *app.RequestContext, p *post.Post) error {
	if err := db.GetDBFromContext(c).Where("id = ? AND deleted = ?", p.ID, false).Updates(p).Error; err != nil {
		return err
	}
	return nil
}

// DeletePost 删除文章（软删除）
func (m *PostMapperImpl) DeletePost(c *app.RequestContext, postID int64) error {
	if err := db.GetDBFromContext(c).Model(&post.Post{}).Where("id = ? AND deleted = ?", postID, false).Update("deleted", true).Error; err != nil {
		return err
	}
	return nil
}
