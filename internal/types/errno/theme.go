// Package errno 主题模块错误码定义
// 创建者：Done-0
// 创建时间：2025-08-05
package errno

import (
	"github.com/Done-0/jank/internal/utils/errorx/code"
)

// 主题模块错误码: 30000 ~ 39999
const (
	ErrThemeSwitchFailed     = 30001 // 切换主题失败
	ErrThemeGetFailed        = 30002 // 获取主题失败
	ErrThemeListFailed       = 30003 // 获取主题列表失败
	ErrThemeNotFound         = 30004 // 主题不存在
	ErrThemeResourceNotFound = 30005 // 主题资源不存在
	ErrThemeConfigInvalid    = 30006 // 主题配置无效
	ErrThemeTypePermissionDenied = 30007 // 主题类型权限被拒绝
)

func init() {
	code.Register(ErrThemeSwitchFailed, "switch theme failed: {theme_id}")
	code.Register(ErrThemeGetFailed, "get theme failed: {msg}")
	code.Register(ErrThemeListFailed, "list themes failed: {msg}")
	code.Register(ErrThemeNotFound, "theme not found: {theme_id}")
	code.Register(ErrThemeTypePermissionDenied, "theme type permission denied: {msg}")
}
