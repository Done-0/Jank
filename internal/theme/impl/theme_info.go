// Package impl 提供主题系统的具体实现
// 创建者：Done-0
// 创建时间：2025-08-09
package impl

// ThemeInfo 主题元数据和运行时信息
type ThemeInfo struct {
	// 基本信息
	ID          string `json:"id"`                    // 主题唯一标识
	Name        string `json:"name,omitempty"`        // 主题显示名称
	Version     string `json:"version,omitempty"`     // 主题版本号
	Author      string `json:"author,omitempty"`      // 主题作者
	Description string `json:"description,omitempty"` // 主题描述
	Repository  string `json:"repository,omitempty"`  // 主题仓库地址
	Preview     string `json:"preview,omitempty"`     // 主题预览图片地址

	// 配置信息
	Type          string `json:"type,omitempty"`            // 主题类型（frontend/console）
	IndexFilePath string `json:"index_file_path,omitempty"` // 首页文件相对路径
	StaticDirPath string `json:"static_dir_path,omitempty"` // 静态资源目录相对路径

	// 运行时信息
	Status   string `json:"status"`              // 当前状态（ready/active/inactive/error）
	LoadedAt int64  `json:"loaded_at,omitempty"` // 加载时间戳
	Path     string `json:"path"`                // 主题文件路径
	IsActive bool   `json:"is_active"`           // 是否为当前激活主题
}
