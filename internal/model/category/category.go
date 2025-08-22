// Package category 提供分类数据模型定义
// 创建者：Done-0
// 创建时间：2025-08-13
package category

import (
	"github.com/Done-0/jank/internal/model/base"
)

// Category 分类模型
type Category struct {
	base.Base
	Name        string `gorm:"type:varchar(100);not null;index" json:"name"`              // 分类名称
	Description string `gorm:"type:varchar(500)" json:"description"`                      // 分类描述（可选）
	ParentID    int64  `gorm:"type:bigint;not null;default:0;index" json:"parent_id"`     // 父分类 ID，0 表示顶级分类
	Sort        int64  `gorm:"type:bigint;not null;default:100;index" json:"sort"`        // 排序权重，数字越大越靠前
	IsActive    bool   `gorm:"type:boolean;not null;default:true;index" json:"is_active"` // 是否启用
}

// TableName 指定表名
// 返回值：
//   - string: 表名
func (Category) TableName() string {
	return "categories"
}
