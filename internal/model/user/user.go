// Package model 提供用户用户数据模型定义
// 创建者：Done-0
// 创建时间：2025-08-05
package user

import "github.com/Done-0/jank/internal/model/base"

// User 用户用户模型
type User struct {
	base.Base
	Email    string `gorm:"type:varchar(64);unique;not null" json:"email"`    // 邮箱，主登录方式
	Password string `gorm:"type:varchar(255);not null" json:"password"`       // 加密密码
	Nickname string `gorm:"type:varchar(64);unique;not null" json:"nickname"` // 昵称
	Avatar   string `gorm:"type:varchar(255);default:null" json:"avatar"`     // 用户头像
	Role     string `gorm:"type:varchar(32);default:'user'" json:"role"`      // 用户角色
}

// TableName 指定表名
// 返回值：
//
//	string: 表名
func (User) TableName() string {
	return "users"
}
