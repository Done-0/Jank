// Package controller 文章控制器
// 创建者：Done-0
// 创建时间：2025-08-13
package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
	"github.com/Done-0/jank/internal/utils/validator"
	"github.com/Done-0/jank/internal/utils/vo"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
)

// PostController 文章控制器
type PostController struct {
	postService service.PostService
}

// NewPostController 创建文章控制器
func NewPostController(postService service.PostService) *PostController {
	return &PostController{
		postService: postService,
	}
}

// GetPost 获取单篇文章
// @Router /api/v1/post/get [get]
func (pc *PostController) GetPost(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetPostRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.postService.GetPost(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPostGetFailed, errorx.KV("id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListPublishedPosts 获取文章列表
// @Router /api/v1/post/list-published [get]
func (pc *PostController) ListPublishedPosts(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListPublishedPostsRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.postService.ListPublishedPosts(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPostListFailed, errorx.KV("msg", "list posts failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListPostsByStatus 根据状态获取文章列表
// @Router /api/v1/post/list-by-status [get]
func (pc *PostController) ListPostsByStatus(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListPostsByStatusRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.postService.ListPostsByStatus(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPostListFailed, errorx.KV("msg", "list posts by status failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Create 创建文章
// @Router /api/v1/post/create [post]
func (pc *PostController) Create(ctx context.Context, c *app.RequestContext) {
	req := new(dto.CreatePostRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.postService.Create(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPostCreateFailed, errorx.KV("title", req.Title))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Update 更新文章
// @Router /api/v1/post/update [post]
func (pc *PostController) Update(ctx context.Context, c *app.RequestContext) {
	req := new(dto.UpdatePostRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.postService.Update(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPostUpdateFailed, errorx.KV("id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Delete 删除文章
// @Router /api/v1/post/delete [post]
func (pc *PostController) Delete(ctx context.Context, c *app.RequestContext) {
	req := new(dto.DeletePostRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.postService.Delete(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPostDeleteFailed, errorx.KV("id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
