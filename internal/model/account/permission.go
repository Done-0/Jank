package model

import "jank.com/jank_blog/internal/model/base"

// Permission 权限模型
type Permission struct {
	base.Base
	Code        string `gorm:"type:varchar(32);unique;not null" json:"code"`       // 权限编码，如 'read', 'write', 'delete'
	Description string `gorm:"type:varchar(255);default: null" json:"description"` // 权限描述
}

func (Permission) TableName() string {
	return "permissions"
}
