// Package base 提供基础模型定义和通用数据库操作方法
// 创建者：Done-0
// 创建时间：2025-08-05
package base

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"

	"github.com/Done-0/jank/internal/utils/snowflake"
)

// Base 包含模型通用字段
type Base struct {
	ID          int64   `gorm:"primaryKey;type:bigint" json:"id"`          // 主键（snowflake）
	GmtCreated  int64   `gorm:"type:bigint" json:"gmt_created"`            // 创建时间
	GmtModified int64   `gorm:"type:bigint" json:"gmt_modified"`           // 更新时间
	Ext         JSONMap `gorm:"type:json" json:"ext"`                      // 扩展字段
	Deleted     bool    `gorm:"type:boolean;default:false" json:"deleted"` // 逻辑删除
}

// JSONMap 处理 JSON 类型字段
type JSONMap map[string]any

// Scan 从数据库读取 JSON 数据
// 参数：
//
//	value: 数据库返回的值
//
// 返回值：
//
//	error: 操作过程中的错误
func (j *JSONMap) Scan(value any) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("data type error, cannot convert to []byte type")
	}
	return json.Unmarshal(bytes, j)
}

// Value 将 JSONMap 转换为 JSON 数据存储到数据库
// 返回值：
//
//	driver.Value: 数据库驱动值
//	error: 操作过程中的错误
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

// BeforeCreate 创建前设置时间戳和 ID
// 参数：
//
//	db: GORM数据库连接
//
// 返回值：
//
//	error: 操作过程中的错误
func (m *Base) BeforeCreate(db *gorm.DB) error {
	now := time.Now().Unix()
	m.GmtCreated = now
	m.GmtModified = now

	// 生成雪花算法 ID
	id, err := snowflake.GenerateID()
	if err != nil {
		return err
	}
	m.ID = id

	if m.Ext == nil {
		m.Ext = make(map[string]any)
	}
	return nil
}

// BeforeUpdate 更新前更新修改时间
// 参数：
//
//	db: GORM数据库连接
//
// 返回值：
//
//	error: 操作过程中的错误
func (m *Base) BeforeUpdate(db *gorm.DB) error {
	m.GmtModified = time.Now().Unix()
	return nil
}
