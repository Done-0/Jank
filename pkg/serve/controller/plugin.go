// Package controller 插件控制器
// 创建者：Done-0
// 创建时间：2025-08-05
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

// PluginController 插件控制器
type PluginController struct {
	pluginService service.PluginService
}

// NewPluginController 创建插件控制器
// 参数：
//
//	pluginService: 插件服务
//
// 返回值：
//
//	*PluginController: 插件控制器
func NewPluginController(pluginService service.PluginService) *PluginController {
	return &PluginController{
		pluginService: pluginService,
	}
}

// RegisterPlugin 注册插件
// @Router /api/v1/plugin/register [post]
func (pc *PluginController) RegisterPlugin(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RegisterPluginRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.pluginService.RegisterPlugin(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPluginRegisterFailed, errorx.KV("plugin_id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// UnregisterPlugin 注销插件
// @Router /api/v1/plugin/unregister [post]
func (pc *PluginController) UnregisterPlugin(ctx context.Context, c *app.RequestContext) {
	req := new(dto.UnregisterPluginRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"), errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.pluginService.UnregisterPlugin(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPluginUnregisterFailed, errorx.KV("plugin_id", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ExecutePlugin 执行插件方法
// @Router /api/v1/plugin/execute [post]
func (pc *PluginController) ExecutePlugin(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ExecutePluginRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.pluginService.ExecutePlugin(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrExecutePluginFailed, errorx.KV("msg", "execute plugin failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetPlugin 获取插件信息
// @Router /api/v1/plugin/get [get]
func (pc *PluginController) GetPlugin(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetPluginRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.pluginService.GetPlugin(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrPluginNotFound, errorx.KV("msg", req.ID))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListPlugins 列举插件
// @Router /api/v1/plugin/list [get]
func (pc *PluginController) ListPlugins(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListPluginsRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := pc.pluginService.ListPlugins(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrListPluginsFailed, errorx.KV("msg", "list plugins failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
