// Package controller RBAC 控制器
// 创建者：Done-0
// 创建时间：2025-08-08
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

// RBACController RBAC控制器
type RBACController struct {
	rbacService service.RBACService
}

// NewRBACController 创建RBAC控制器实例
func NewRBACController(rbacService service.RBACService) *RBACController {
	return &RBACController{
		rbacService: rbacService,
	}
}

// CreatePermission 创建权限
// @Router /api/v1/rbac/create-permission [post]
func (rc *RBACController) CreatePermission(ctx context.Context, c *app.RequestContext) {
	req := new(dto.CreatePermissionRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.CreatePermission(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "create permission failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// DeletePermission 删除权限
// @Router /api/v1/rbac/delete-permission [post]
func (rc *RBACController) DeletePermission(ctx context.Context, c *app.RequestContext) {
	req := new(dto.DeletePermissionRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.DeletePermission(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "delete permission failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListPermissions 获取所有权限
// @Router /api/v1/rbac/list-permissions [get]
func (rc *RBACController) ListPermissions(ctx context.Context, c *app.RequestContext) {
	response, err := rc.rbacService.ListPermissions(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "list permissions failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// AssignPermission 为角色分配权限
// @Router /api/v1/rbac/assign-permission [post]
func (rc *RBACController) AssignPermission(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AssignPermissionRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.AssignPermission(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "assign permission failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RevokePermission 撤销角色权限
// @Router /api/v1/rbac/revoke-permission [post]
func (rc *RBACController) RevokePermission(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RevokePermissionRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.RevokePermission(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "revoke permission failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// CreateRole 创建角色
// @Router /api/v1/rbac/create-role [post]
func (rc *RBACController) CreateRole(ctx context.Context, c *app.RequestContext) {
	req := new(dto.CreateRoleRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.CreateRole(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "create role failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// DeleteRole 删除角色
// @Router /api/v1/rbac/delete-role [post]
func (rc *RBACController) DeleteRole(ctx context.Context, c *app.RequestContext) {
	req := new(dto.DeleteRoleRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.DeleteRole(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "delete role failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListRoles 获取所有角色
// @Router /api/v1/rbac/list-roles [get]
func (rc *RBACController) ListRoles(ctx context.Context, c *app.RequestContext) {
	response, err := rc.rbacService.ListRoles(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "list roles failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetRolePermissions 获取角色权限
// @Router /api/v1/rbac/get-role-permissions [post]
func (rc *RBACController) GetRolePermissions(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetRolePermissionsRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.GetRolePermissions(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "get role permissions failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// AssignRole 为用户分配角色
// @Router /api/v1/rbac/assign-role [post]
func (rc *RBACController) AssignRole(ctx context.Context, c *app.RequestContext) {
	req := new(dto.AssignRoleRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.AssignRole(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "assign role failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RevokeRole 撤销用户角色
// @Router /api/v1/rbac/revoke-role [post]
func (rc *RBACController) RevokeRole(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RevokeRoleRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.RevokeRole(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "revoke role failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetUserRoles 获取用户角色
// @Router /api/v1/rbac/get-user-roles [post]
func (rc *RBACController) GetUserRoles(ctx context.Context, c *app.RequestContext) {
	req := new(dto.GetUserRolesRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.GetUserRoles(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "get user roles failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// CheckPermission 权限检查
// @Router /api/v1/rbac/check-permission [post]
func (rc *RBACController) CheckPermission(ctx context.Context, c *app.RequestContext) {
	req := new(dto.CheckPermissionRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := rc.rbacService.CheckPermission(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "check permission failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
