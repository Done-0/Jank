// Package controller 主题控制器
// 创建者：Done-0
// 创建时间：2025-08-05
package controller

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
	"github.com/Done-0/jank/internal/utils/validator"
	"github.com/Done-0/jank/internal/utils/vo"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
)

// ThemeController 主题控制器
type ThemeController struct {
	themeService service.ThemeService
}

// NewThemeController 创建主题控制器
func NewThemeController(themeService service.ThemeService) *ThemeController {
	return &ThemeController{
		themeService: themeService,
	}
}

// SwitchTheme 切换主题
// @Router /api/v1/theme/switch [post]
func (tc *ThemeController) SwitchTheme(ctx context.Context, c *app.RequestContext) {
	req := new(dto.SwitchThemeRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := tc.themeService.SwitchTheme(c, req)
	if err != nil {
		// 检查是否是主题类型权限控制错误
		if strings.Contains(err.Error(), "can only switch to") ||
			strings.Contains(err.Error(), "theme type mismatch") ||
			strings.Contains(err.Error(), "permission denied") {
			c.JSON(consts.StatusForbidden, vo.Fail(c, err, errorx.New(errno.ErrThemeTypePermissionDenied, errorx.KV("msg", err.Error()))))
			return
		}

		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrThemeSwitchFailed, errorx.KV("theme_id", req.ID))))
		return
	}

	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetActiveTheme 获取当前激活主题
// @Router /api/v1/theme/get [get]
func (tc *ThemeController) GetActiveTheme(ctx context.Context, c *app.RequestContext) {
	response, err := tc.themeService.GetActiveTheme(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrThemeGetFailed, errorx.KV("msg", "get active theme failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListThemes 列举主题
// @Router /api/v1/theme/list [get]
func (tc *ThemeController) ListThemes(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListThemesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := tc.themeService.ListThemes(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrThemeListFailed, errorx.KV("msg", "list themes failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ServeHomePage 提供主题首页
// @Router / [get]
func (tc *ThemeController) ServeHomePage(ctx context.Context, c *app.RequestContext) {
	homePagePath, err := tc.themeService.ServeHomePage(c)
	if err != nil {
		c.AbortWithStatus(consts.StatusInternalServerError)
		return
	}

	c.File(homePagePath)
}

// ServeStaticResource 提供静态资源文件
// @Router /* [get]
func (tc *ThemeController) ServeStaticResource(ctx context.Context, c *app.RequestContext) {
	staticResourcePath, err := tc.themeService.ServeStaticResource(c, string(c.Path()))
	if err != nil {
		c.AbortWithStatus(consts.StatusNotFound)
		return
	}

	c.File(staticResourcePath)
}
