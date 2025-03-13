package model

import (
	"jank.com/jank_blog/internal/model/base"
)

// PluginInfo 插件基本信息模型
type PluginInfo struct {
	base.Base
	Name          string         `gorm:"type:varchar(128);not null" json:"name"`   // Name 插件名称
	Version       string         `gorm:"type:varchar(64);not null" json:"version"` // Version 插件版本
	Description   string         `gorm:"type:varchar(2048)" json:"description"`    // Description 插件描述
	Author        string         `gorm:"type:varchar(128)" json:"author"`          // Author 插件作者
	Tags          []string       `gorm:"type:json" json:"tags"`                    // Tags 插件标签
	Requires      []string       `gorm:"type:json" json:"requires"`                // Requires 插件依赖
	DefaultConfig map[string]any `gorm:"type:json" json:"defaultConfig"`           // DefaultConfig 插件默认配置
}

func (PluginInfo) TableName() string {
	return "plugin_infos"
}
