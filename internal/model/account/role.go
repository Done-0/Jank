package model

import "jank.com/jank_blog/internal/model/base"

// Role 角色模型
type Role struct {
	base.Base
	Code        string `gorm:"type:varchar(32);unique;not null" json:"code"`       // 角色编码，如 'admin', 'user'
	Description string `gorm:"type:varchar(255);default: null" json:"description"` // 角色描述
}

func (Role) TableName() string {
	return "roles"
}
