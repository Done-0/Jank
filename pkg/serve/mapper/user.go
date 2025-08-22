// Package mapper 提供用户相关的数据访问接口
// 创建者：Done-0
// 创建时间：2025-08-10
package mapper

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/model/user"
)

// UserMapper 用户数据访问接口
type UserMapper interface {
	// 基础用户操作
	GetUserByEmail(c *app.RequestContext, email string) (*user.User, error)       // 根据邮箱获取用户
	GetUserByID(c *app.RequestContext, userID int64) (*user.User, error)          // 根据 ID 获取用户
	GetUserByNickname(c *app.RequestContext, nickname string) (*user.User, error) // 根据昵称获取用户

	RegisterUser(c *app.RequestContext, user *user.User) error // 注册用户
	UpdateUser(c *app.RequestContext, user *user.User) error   // 更新用户信息

	// 用户管理操作
	ListUsers(c *app.RequestContext, pageNo, pageSize int64, keyword, role string) ([]*user.User, int64, error) // 获取用户列表
	DeleteUser(c *app.RequestContext, userID string) error                                                      // 删除用户
}
