package model

import "jank.com/jank_blog/internal/model/base"

// AccountRole 用户和角色关联表
type AccountRole struct {
	base.Base
	AccountID int64 `gorm:"index" json:"account_id"` // 用户ID
	RoleID    int64 `gorm:"index" json:"role_id"`    // 角色ID
}

func (AccountRole) TableName() string {
	return "account_roles"
}
