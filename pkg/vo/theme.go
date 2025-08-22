// Package vo 主题相关值对象
package vo

// SwitchThemeResponse 切换主题响应
type SwitchThemeResponse struct {
	Message string `json:"message"` // 切换结果消息
}

// GetThemeResponse 主题信息
type GetThemeResponse struct {
	// 基本信息
	ID          string `json:"id"`                    // 主题唯一标识
	Name        string `json:"name"`                  // 主题显示名称
	Version     string `json:"version"`               // 主题版本号
	Author      string `json:"author"`                // 主题作者
	Description string `json:"description,omitempty"` // 主题描述
	Repository  string `json:"repository,omitempty"`  // 主题仓库地址
	Preview     string `json:"preview,omitempty"`     // 主题预览图片地址

	// 配置信息
	Type          string `json:"type"`                      // 主题类型（frontend/console）
	IndexFilePath string `json:"index_file_path,omitempty"` // 首页文件相对路径
	StaticDirPath string `json:"static_dir_path,omitempty"` // 静态资源目录相对路径

	// 运行时信息
	Status   string `json:"status"`              // 当前状态（ready/active/inactive/error）
	LoadedAt int64  `json:"loaded_at,omitempty"` // 加载时间戳
	Path     string `json:"path"`                // 主题文件路径
	IsActive bool   `json:"is_active"`           // 是否为当前激活主题
}

// GetActiveThemeResponse 获取当前激活主题响应
type GetActiveThemeResponse struct {
	Theme *GetThemeResponse `json:"theme"` // 当前激活的主题信息
}

// ListThemesResponse 列举主题响应
type ListThemesResponse struct {
	Total    int64              `json:"total"`     // 总数
	PageNo   int64              `json:"page_no"`   // 当前页码
	PageSize int64              `json:"page_size"` // 每页大小
	List     []GetThemeResponse `json:"list"`      // 主题列表
}
