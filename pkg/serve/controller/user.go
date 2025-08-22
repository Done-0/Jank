// Package controller 用户控制器
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

// UserController 用户控制器
type UserController struct {
	userService service.UserService
}

// NewUserController 创建用户控制器
func NewUserController(userService service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

// Register 用户注册
// @Router /api/v1/user/register [post]
func (uc *UserController) Register(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RegisterRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.Register(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserRegisterFailed, errorx.KV("email", req.Email))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Login 用户登录
// @Router /api/v1/user/login [post]
func (uc *UserController) Login(ctx context.Context, c *app.RequestContext) {
	req := new(dto.LoginRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.Login(c, req)
	if err != nil {
		c.JSON(consts.StatusUnauthorized, vo.Fail(c, err, errorx.New(errno.ErrUserLoginFailed, errorx.KV("email", req.Email))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Logout 用户登出
// @Router /api/v1/user/logout [post]
func (uc *UserController) Logout(ctx context.Context, c *app.RequestContext) {
	response, err := uc.userService.Logout(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserLogoutFailed, errorx.KV("msg", "logout failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// RefreshToken 刷新token
// @Router /api/v1/user/refresh-token [post]
func (uc *UserController) RefreshToken(ctx context.Context, c *app.RequestContext) {
	req := new(dto.RefreshTokenRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.RefreshToken(c, req)
	if err != nil {
		c.JSON(consts.StatusUnauthorized, vo.Fail(c, err, errorx.New(errno.ErrUserRefreshTokenFailed, errorx.KV("msg", "refresh token failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// GetProfile 获取用户资料
// @Router /api/v1/user/profile [get]
func (uc *UserController) GetProfile(ctx context.Context, c *app.RequestContext) {
	response, err := uc.userService.GetProfile(c)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserGetProfileFailed, errorx.KV("msg", "get profile failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// Update 更新用户信息
// @Router /api/v1/user/update [post]
func (uc *UserController) Update(ctx context.Context, c *app.RequestContext) {
	req := new(dto.UpdateRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.Update(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "update user failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ResetPassword 重置密码
// @Router /api/v1/user/reset-password [post]
func (uc *UserController) ResetPassword(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ResetPasswordRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.ResetPassword(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserResetPasswordFailed, errorx.KV("msg", "reset password failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// ListUsers 列举用户（管理员功能）
// @Router /api/v1/user/list [get]
func (uc *UserController) ListUsers(ctx context.Context, c *app.RequestContext) {
	req := new(dto.ListUsersRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation"), errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.ListUsers(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrUserListFailed, errorx.KV("msg", "list users failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}

// UpdateUserRole 管理员更新用户角色
// @Router /api/v1/user/role [post]
func (uc *UserController) UpdateUserRole(ctx context.Context, c *app.RequestContext) {
	req := new(dto.UpdateUserRoleRequest)
	if err := c.BindJSON(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind JSON failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	response, err := uc.userService.UpdateUserRole(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "update user role failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, response))
}
