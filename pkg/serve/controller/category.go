// Package controller 分类控制器
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

// CategoryController 分类控制器
type CategoryController struct {
	categoryService service.CategoryService
}

// NewCategoryController 创建分类控制器
func NewCategoryController(categoryService service.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: categoryService,
	}
}

// GetCategory 获取单个分类
// @Router /api/v1/category/get [get]
func (cc *CategoryController) GetCategory(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetCategoryRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := cc.categoryService.GetCategory(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrCategoryGetFailed, errorx.KV("id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListCategories 获取分类列表
// @Router /api/v1/category/list [get]
func (cc *CategoryController) ListCategories(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListCategoriesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := cc.categoryService.ListCategories(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrCategoryListFailed, errorx.KV("msg", "list categories failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Create 创建分类
// @Router /api/v1/category/create [post]
func (cc *CategoryController) Create(ctx context.Context, c *app.RequestContext) {
	req := new(dto.CreateCategoryRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := cc.categoryService.Create(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrCategoryCreateFailed, errorx.KV("name", req.Name))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Update 更新分类
// @Router /api/v1/category/update [post]
func (cc *CategoryController) Update(ctx context.Context, c *app.RequestContext) {
	req := new(dto.UpdateCategoryRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := cc.categoryService.Update(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrCategoryUpdateFailed, errorx.KV("id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Delete 删除分类
// @Router /api/v1/category/delete [post]
func (cc *CategoryController) Delete(ctx context.Context, c *app.RequestContext) {
	req := new(dto.DeleteCategoryRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := cc.categoryService.Delete(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrCategoryDeleteFailed, errorx.KV("id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
