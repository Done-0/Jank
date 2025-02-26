package model

import (
	"database/sql/driver"
	"encoding/json"
	"jank.com/jank_blog/internal/model/base"
)

// Plugin 插件模型
type Plugin struct {
	base.Base
	Name         string       `gorm:"size:100;not null" json:"name"`           // 插件名称
	Version      string       `gorm:"size:50;not null" json:"version"`         // 插件版本
	Type         PluginType   `gorm:"size:50;not null" json:"type"`            // 插件类型（主插件/依赖/日志/执行记录）
	Status       PluginStatus `gorm:"type:varchar(20);not null" json:"status"` // 插件状态（已安装/激活/未激活/错误）
	IsEnabled    bool         `gorm:"default:false" json:"is_enabled"`         // 是否启用（true/false）
	Data         base.JSONMap `gorm:"type:json" json:"data"`                   // 通用数据（插件相关信息）
	Dependencies Int64Array   `gorm:"type:json" json:"dependencies"`           // 依赖插件ID列表
}

// PluginType 插件类型枚举
type PluginType string

const (
	PluginTypeMain       PluginType = "main"       // 主插件
	PluginTypeDependency PluginType = "dependency" // 依赖插件
	PluginTypeLog        PluginType = "log"        // 日志
	PluginTypeExecution  PluginType = "execution"  // 执行记录
)

// PluginStatus 插件状态枚举
type PluginStatus string

const (
	PluginStatusInstalled PluginStatus = "installed" // 已安装
	PluginStatusActive    PluginStatus = "active"    // 激活
	PluginStatusInactive  PluginStatus = "inactive"  // 未激活
	PluginStatusError     PluginStatus = "error"     // 错误
)

// Int64Array int64数组类型
type Int64Array []int64

// Value 实现 driver.Valuer 接口
func (a Int64Array) Value() (driver.Value, error) {
	if a == nil {
		return nil, nil
	}
	return json.Marshal(a)
}

// Scan 实现 sql.Scanner 接口
func (a *Int64Array) Scan(value interface{}) error {
	if value == nil {
		*a = nil
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return nil
	}
	return json.Unmarshal(bytes, a)
}

// NewPlugin 创建新的插件实例
func NewPlugin(name, version string, pluginType PluginType) *Plugin {
	return &Plugin{
		Name:         name,
		Version:      version,
		Type:         pluginType,
		Status:       PluginStatusInstalled,
		IsEnabled:    false,
		Data:         make(base.JSONMap),
		Dependencies: make(Int64Array, 0),
	}
}

// AddDependency 添加依赖插件
func (p *Plugin) AddDependency(dependencyID int64) {
	p.Dependencies = append(p.Dependencies, dependencyID)
}

// RemoveDependency 移除依赖插件
func (p *Plugin) RemoveDependency(dependencyID int64) {
	for i, dep := range p.Dependencies {
		if dep == dependencyID {
			p.Dependencies = append(p.Dependencies[:i], p.Dependencies[i+1:]...)
			break
		}
	}
}

// HasDependency 检查是否有特定的依赖插件
func (p *Plugin) HasDependency(dependencyID int64) bool {
	for _, dep := range p.Dependencies {
		if dep == dependencyID {
			return true
		}
	}
	return false
}

// SetData 设置插件数据
func (p *Plugin) SetData(key string, value interface{}) {
	if p.Data == nil {
		p.Data = make(base.JSONMap)
	}
	p.Data[key] = value
}

// GetData 获取插件数据
func (p *Plugin) GetData(key string) (interface{}, bool) {
	if p.Data == nil {
		return nil, false
	}
	value, exists := p.Data[key]
	return value, exists
}

// GetStringArray 从Data中获取字符串数组
func (p *Plugin) GetStringArray(key string) ([]string, bool) {
	value, exists := p.GetData(key)
	if !exists {
		return nil, false
	}

	if strArray, ok := value.([]string); ok {
		return strArray, true
	}

	if jsonStr, ok := value.(string); ok {
		var strArray []string
		err := json.Unmarshal([]byte(jsonStr), &strArray)
		if err == nil {
			return strArray, true
		}
	}

	return nil, false
}
