// Package base 提供基础模型定义和通用数据库操作方法
// 创建者：Done-0
// 创建时间：2025-05-10
package base

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"log"
	"time"

	"gorm.io/gorm"

	"jank.com/jank_blog/internal/utils"
)

// Base 包含通用字段
type Base struct {
	ID          int64   `gorm:"primaryKey;type:bigint" json:"id"`          // 主键（雪花算法）
	GmtCreate   int64   `gorm:"type:bigint" json:"gmt_create"`             // 创建时间
	GmtModified int64   `gorm:"type:bigint" json:"gmt_modified"`           // 更新时间
	Ext         JSONMap `gorm:"type:json" json:"ext"`                      // 扩展字段
	Deleted     bool    `gorm:"type:boolean;default:false" json:"deleted"` // 逻辑删除
}

// JSONMap 处理 json 类型字段
type JSONMap map[string]interface{}

// Scan 从数据库读取 json 数据
// 参数：
//   - value: 数据库返回的值
//
// 返回值：
//   - error: 操作过程中的错误
func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型错误，无法转换为 []byte 类型")
	}
	return json.Unmarshal(bytes, j)
}

// Value 将 JSONMap 转换为 json 数据存储到数据库
// 返回值：
//   - driver.Value: 数据库驱动值
//   - error: 操作过程中的错误
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

// BeforeCreate 创建前操作，设置时间戳等
// 参数：
//   - db: GORM数据库连接
//
// 返回值：
//   - error: 操作过程中的错误
func (m *Base) BeforeCreate(db *gorm.DB) (err error) {
	currentTime := time.Now().Unix()
	m.GmtCreate = currentTime
	m.GmtModified = currentTime
	m.Deleted = false

	// 使用雪花算法生成ID
	id, err := utils.GenerateID()
	if err != nil {
		log.Printf("生成雪花ID时出错: %v", err)
	}
	m.ID = id

	if m.Ext == nil {
		m.Ext = make(map[string]interface{})
	}
	return nil
}

// BeforeUpdate 更新前操作，更新修改时间
// 参数：
//   - db: GORM数据库连接
//
// 返回值：
//   - error: 操作过程中的错误
func (m *Base) BeforeUpdate(db *gorm.DB) (err error) {
	m.GmtModified = time.Now().Unix()
	return nil
}
