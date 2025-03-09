package base

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"

	"gorm.io/gorm"
)

// Base 包含通用字段
type Base struct {
	ID          int64   `gorm:"primaryKey;autoIncrement" json:"id"`        // 主键
	GmtCreate   int64   `gorm:"type:bigint" json:"gmt_create"`             // 创建时间
	GmtModified int64   `gorm:"type:bigint" json:"gmt_modified"`           // 更新时间
	Ext         JSONMap `gorm:"type:json" json:"ext"`                      // 扩展字段
	Deleted     bool    `gorm:"type:boolean;default:false" json:"deleted"` // 逻辑删除
}

// JSONMap 处理 json 类型字段
type JSONMap map[string]interface{}

// Scan 从数据库读取 json 数据
func (j *JSONMap) Scan(value interface{}) error {
	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("数据类型错误，无法转换为 []byte 类型")
	}
	return json.Unmarshal(bytes, j)
}

// Value 将 JSONMap 转换为 json 数据存储到数据库
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return "{}", nil
	}
	return json.Marshal(j)
}

// BeforeCreate 创建前操作，设置时间戳等
func (m *Base) BeforeCreate(db *gorm.DB) (err error) {
	currentTime := time.Now().Unix()
	m.GmtCreate = currentTime
	m.GmtModified = currentTime
	m.Deleted = false
	if m.Ext == nil {
		m.Ext = make(map[string]interface{})
	}
	return nil
}

// BeforeUpdate 更新前操作，更新修改时间
func (m *Base) BeforeUpdate(db *gorm.DB) (err error) {
	m.GmtModified = time.Now().Unix()
	return nil
}
