// Package rbac 提供RBAC权限相关的数据模型定义
// 创建者：Done-0
// 创建时间：2025-08-05
package rbac

import (
	"gorm.io/gorm"

	"github.com/Done-0/jank/internal/model/base"
)

// Policy RBAC 策略模型
type Policy struct {
	base.Base
	Ptype string `gorm:"type:varchar(10);not null" json:"ptype"` // 策略类型: p(权限)或g(角色)
	V0    string `gorm:"type:varchar(100);not null" json:"v0"`   // 主体(角色名)
	V1    string `gorm:"type:varchar(100);not null" json:"v1"`   // 资源路径或继承的角色
	V2    string `gorm:"type:varchar(100)" json:"v2"`            // 操作方法(GET, POST等)
	V3    string `gorm:"type:varchar(100)" json:"v3"`            // 保留字段
	V4    string `gorm:"type:varchar(100)" json:"v4"`            // 保留字段
	V5    string `gorm:"type:varchar(100)" json:"v5"`            // 保留字段
}

// TableName 指定表名
// 返回值：
//
//	string: 表名
func (Policy) TableName() string {
	return "casbin_rule"
}

// BeforeCreate 创建前钩子，确保基础字段正确生成
func (p *Policy) BeforeCreate(db *gorm.DB) error {
	return p.Base.BeforeCreate(db)
}

// BeforeUpdate 更新前钩子，确保修改时间正确更新
func (p *Policy) BeforeUpdate(db *gorm.DB) error {
	return p.Base.BeforeUpdate(db)
}
