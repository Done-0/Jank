package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// Account 用户账户模型
type Account struct {
	base.BaseModel
	Phone    string `gorm:"type:varchar(32);unique;null" json:"phone"`                 // 手机号，次登录方式
	Email    string `gorm:"type:varchar(64);unique;not null" json:"email"`             // 邮箱，主登录方式
	Password string `gorm:"type:varchar(255);not null" json:"password"`                // 加密密码
	Nickname string `gorm:"type:varchar(64);not null" json:"nickname"`                 // 昵称
	RoleCode string `gorm:"type:varchar(32);not null;default:'user'" json:"role_code"` // 角色编码，默认为 user
	Avatar   string `gorm:"type:varchar(255);default:null" json:"avatar"`              // 用户头像
}

func (Account) TableName() string {
	return "accounts"
}
