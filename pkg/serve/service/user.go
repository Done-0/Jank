package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/vo"
)

// UserService 用户服务接口
type UserService interface {
	Register(c *app.RequestContext, req *dto.RegisterRequest) (*vo.RegisterResponse, error)                   // 注册用户
	Login(c *app.RequestContext, req *dto.LoginRequest) (*vo.LoginResponse, error)                            // 用户登录
	Logout(c *app.RequestContext) (*vo.LogoutResponse, error)                                                 // 注销登录
	RefreshToken(c *app.RequestContext, req *dto.RefreshTokenRequest) (*vo.RefreshTokenResponse, error)       // 刷新token
	GetProfile(c *app.RequestContext) (*vo.GetProfileResponse, error)                                         // 获取用户
	Update(c *app.RequestContext, req *dto.UpdateRequest) (*vo.UpdateResponse, error)                         // 更新用户
	ResetPassword(c *app.RequestContext, req *dto.ResetPasswordRequest) (*vo.ResetPasswordResponse, error)    // 重置密码
	ListUsers(c *app.RequestContext, req *dto.ListUsersRequest) (*vo.ListUsersResponse, error)                // 获取用户列表
	UpdateUserRole(c *app.RequestContext, req *dto.UpdateUserRoleRequest) (*vo.UpdateUserRoleResponse, error) // 管理员更新用户角色
}
