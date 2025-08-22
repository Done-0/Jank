// Package impl 文章服务实现
// 创建者：Done-0
// 创建时间：2025-08-13
package impl

import (
	"fmt"
	"strconv"
	"time"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/post"
	"github.com/Done-0/jank/internal/types/consts"
	"github.com/Done-0/jank/internal/utils/logger"
	"github.com/Done-0/jank/internal/utils/markdown"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/mapper"
	"github.com/Done-0/jank/pkg/serve/service"
	"github.com/Done-0/jank/pkg/vo"
)

// PostServiceImpl 文章服务实现
type PostServiceImpl struct {
	postMapper     mapper.PostMapper
	categoryMapper mapper.CategoryMapper
}

// NewPostService 创建文章服务实例
func NewPostService(postMapperImpl mapper.PostMapper, categoryMapperImpl mapper.CategoryMapper) service.PostService {
	return &PostServiceImpl{
		postMapper:     postMapperImpl,
		categoryMapper: categoryMapperImpl,
	}
}

// GetPost 获取单篇文章
func (ps *PostServiceImpl) GetPost(c *app.RequestContext, req *dto.GetPostRequest) (*vo.GetPostResponse, error) {
	postID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid post ID format: %s", req.ID)
		return nil, fmt.Errorf("invalid post ID format: %w", err)
	}

	post, err := ps.postMapper.GetPostByID(c, postID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get post with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to get post: %w", err)
	}

	var categoryIDStr, categoryName string
	if post.CategoryID != nil {
		if category, err := ps.categoryMapper.GetCategoryByID(c, *post.CategoryID); err == nil && category.IsActive {
			categoryIDStr = strconv.FormatInt(*post.CategoryID, 10)
			categoryName = category.Name
		}
	}

	return &vo.GetPostResponse{
		ID:           strconv.FormatInt(post.ID, 10),
		Title:        post.Title,
		Description:  post.Description,
		Image:        post.Image,
		Status:       post.Status,
		CategoryID:   categoryIDStr,
		CategoryName: categoryName,
		Markdown:     post.Markdown,
		HTML:         post.HTML,
		CreatedAt:    time.Unix(post.GmtCreated, 0).Format("2006-01-02 15:04:05"),
		UpdatedAt:    time.Unix(post.GmtModified, 0).Format("2006-01-02 15:04:05"),
	}, nil
}

// ListPublishedPosts 获取已发布文章列表
func (ps *PostServiceImpl) ListPublishedPosts(c *app.RequestContext, req *dto.ListPublishedPostsRequest) (*vo.ListPostsResponse, error) {
	posts, total, err := ps.postMapper.ListPublishedPosts(c, req.PageNo, req.PageSize, req.CategoryID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list posts: %v", err)
		return nil, fmt.Errorf("failed to list posts: %w", err)
	}

	postItems := make([]*vo.PostItem, 0, len(posts))
	for _, post := range posts {
		var categoryIDStr, categoryName string
		if post.CategoryID != nil {
			if category, err := ps.categoryMapper.GetCategoryByID(c, *post.CategoryID); err == nil && category.IsActive {
				categoryIDStr = strconv.FormatInt(*post.CategoryID, 10)
				categoryName = category.Name
			}
		}

		postItems = append(postItems, &vo.PostItem{
			ID:           strconv.FormatInt(post.ID, 10),
			Title:        post.Title,
			Description:  post.Description,
			Image:        post.Image,
			Status:       post.Status,
			CategoryID:   categoryIDStr,
			CategoryName: categoryName,
			CreatedAt:    time.Unix(post.GmtCreated, 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:    time.Unix(post.GmtModified, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return &vo.ListPostsResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     postItems,
	}, nil
}

// ListPostsByStatus 根据状态获取文章列表
func (ps *PostServiceImpl) ListPostsByStatus(c *app.RequestContext, req *dto.ListPostsByStatusRequest) (*vo.ListPostsResponse, error) {
	posts, total, err := ps.postMapper.ListPostsByStatus(c, req.PageNo, req.PageSize, req.Status, req.CategoryID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to list posts by status: %v", err)
		return nil, fmt.Errorf("failed to list posts by status: %w", err)
	}

	postItems := make([]*vo.PostItem, 0, len(posts))
	for _, post := range posts {
		var categoryIDStr, categoryName string
		if post.CategoryID != nil {
			if category, err := ps.categoryMapper.GetCategoryByID(c, *post.CategoryID); err == nil && category.IsActive {
				categoryIDStr = strconv.FormatInt(*post.CategoryID, 10)
				categoryName = category.Name
			}
		}

		postItems = append(postItems, &vo.PostItem{
			ID:           strconv.FormatInt(post.ID, 10),
			Title:        post.Title,
			Description:  post.Description,
			Image:        post.Image,
			Status:       post.Status,
			CategoryID:   categoryIDStr,
			CategoryName: categoryName,
			CreatedAt:    time.Unix(post.GmtCreated, 0).Format("2006-01-02 15:04:05"),
			UpdatedAt:    time.Unix(post.GmtModified, 0).Format("2006-01-02 15:04:05"),
		})
	}

	return &vo.ListPostsResponse{
		Total:    total,
		PageNo:   req.PageNo,
		PageSize: req.PageSize,
		List:     postItems,
	}, nil
}

// Create 创建文章
func (ps *PostServiceImpl) Create(c *app.RequestContext, req *dto.CreatePostRequest) (*vo.CreatePostResponse, error) {
	status := req.Status
	if status == "" {
		status = consts.PostStatusDraft
	}

	var htmlContent string
	if req.Markdown != "" {
		html, err := markdown.RenderMarkdown([]byte(req.Markdown))
		if err != nil {
			logger.BizLogger(c).Errorf("failed to render markdown for post '%s': %v", req.Title, err)
			return nil, fmt.Errorf("failed to render markdown: %w", err)
		}
		htmlContent = html
	}

	var categoryID *int64
	if req.CategoryID != "" {
		parsedCategoryID, err := strconv.ParseInt(req.CategoryID, 10, 64)
		if err != nil {
			logger.BizLogger(c).Errorf("invalid category ID format: %s", req.CategoryID)
			return nil, fmt.Errorf("invalid category ID format: %w", err)
		}

		_, err = ps.categoryMapper.GetCategoryByID(c, parsedCategoryID)
		if err != nil {
			logger.BizLogger(c).Errorf("category with ID %d does not exist: %v", parsedCategoryID, err)
			return nil, fmt.Errorf("category with ID %d does not exist", parsedCategoryID)
		}

		categoryID = &parsedCategoryID
	}

	post := &post.Post{
		Title:       req.Title,
		Description: req.Description,
		Image:       req.Image,
		Status:      status,
		CategoryID:  categoryID,
		Markdown:    req.Markdown,
		HTML:        htmlContent,
	}

	if err := ps.postMapper.CreatePost(c, post); err != nil {
		logger.BizLogger(c).Errorf("failed to create post '%s': %v", req.Title, err)
		return nil, fmt.Errorf("failed to create post: %w", err)
	}

	logger.BizLogger(c).Infof("post created successfully with ID: %d", post.ID)

	var categoryIDStr, categoryName string
	if post.CategoryID != nil {
		if category, err := ps.categoryMapper.GetCategoryByID(c, *post.CategoryID); err == nil && category.IsActive {
			categoryIDStr = strconv.FormatInt(*post.CategoryID, 10)
			categoryName = category.Name
		}
	}

	return &vo.CreatePostResponse{
		ID:           strconv.FormatInt(post.ID, 10),
		Title:        post.Title,
		Description:  post.Description,
		Image:        post.Image,
		Status:       post.Status,
		CategoryID:   categoryIDStr,
		CategoryName: categoryName,
		Markdown:     post.Markdown,
		Message:      "Post created successfully",
	}, nil
}

// Update 更新文章
func (ps *PostServiceImpl) Update(c *app.RequestContext, req *dto.UpdatePostRequest) (*vo.UpdatePostResponse, error) {
	postID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid post ID format: %s", req.ID)
		return nil, fmt.Errorf("invalid post ID format: %w", err)
	}

	existingPost, err := ps.postMapper.GetPostByID(c, postID)
	if err != nil {
		logger.BizLogger(c).Errorf("failed to get existing post with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to get existing post: %w", err)
	}

	// 更新字段（只更新非空字段）
	if req.Title != "" {
		existingPost.Title = req.Title
	}
	if req.Description != "" {
		existingPost.Description = req.Description
	}
	if req.Image != "" {
		existingPost.Image = req.Image
	}
	if req.Status != "" {
		existingPost.Status = req.Status
	}
	if req.Markdown != "" {
		existingPost.Markdown = req.Markdown
		html, err := markdown.RenderMarkdown([]byte(req.Markdown))
		if err != nil {
			logger.BizLogger(c).Errorf("failed to render markdown for post ID %s: %v", req.ID, err)
			return nil, fmt.Errorf("failed to render markdown: %w", err)
		}
		existingPost.HTML = html
	}
	if req.CategoryID != "" {
		parsedCategoryID, err := strconv.ParseInt(req.CategoryID, 10, 64)
		if err != nil {
			logger.BizLogger(c).Errorf("invalid category ID format: %s", req.CategoryID)
			return nil, fmt.Errorf("invalid category ID format: %w", err)
		}

		_, err = ps.categoryMapper.GetCategoryByID(c, parsedCategoryID)
		if err != nil {
			logger.BizLogger(c).Errorf("category with ID %d does not exist: %v", parsedCategoryID, err)
			return nil, fmt.Errorf("category with ID %d does not exist", parsedCategoryID)
		}

		existingPost.CategoryID = &parsedCategoryID
	}

	if err := ps.postMapper.UpdatePost(c, existingPost); err != nil {
		logger.BizLogger(c).Errorf("failed to update post with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to update post: %w", err)
	}

	logger.BizLogger(c).Infof("post updated successfully with ID: %s", req.ID)

	var categoryIDStr, categoryName string
	if existingPost.CategoryID != nil {
		if category, err := ps.categoryMapper.GetCategoryByID(c, *existingPost.CategoryID); err == nil && category.IsActive {
			categoryIDStr = strconv.FormatInt(*existingPost.CategoryID, 10)
			categoryName = category.Name
		}
	}

	return &vo.UpdatePostResponse{
		ID:           strconv.FormatInt(existingPost.ID, 10),
		Title:        existingPost.Title,
		Description:  existingPost.Description,
		Image:        existingPost.Image,
		Status:       existingPost.Status,
		CategoryID:   categoryIDStr,
		CategoryName: categoryName,
		Markdown:     existingPost.Markdown,
		Message:      "Post updated successfully",
	}, nil
}

// Delete 删除文章
func (ps *PostServiceImpl) Delete(c *app.RequestContext, req *dto.DeletePostRequest) (*vo.DeletePostResponse, error) {
	postID, err := strconv.ParseInt(req.ID, 10, 64)
	if err != nil {
		logger.BizLogger(c).Errorf("invalid post ID format: %s", req.ID)
		return nil, fmt.Errorf("invalid post ID format: %w", err)
	}

	_, err = ps.postMapper.GetPostByID(c, postID)
	if err != nil {
		logger.BizLogger(c).Errorf("post with ID %s not found: %v", req.ID, err)
		return nil, fmt.Errorf("post not found: %w", err)
	}

	if err := ps.postMapper.DeletePost(c, postID); err != nil {
		logger.BizLogger(c).Errorf("failed to delete post with ID %s: %v", req.ID, err)
		return nil, fmt.Errorf("failed to delete post: %w", err)
	}

	logger.BizLogger(c).Infof("post deleted successfully with ID: %s", req.ID)

	return &vo.DeletePostResponse{
		Message: "Post deleted successfully",
	}, nil
}
