// Package consts 提供文章相关常量定义
// 创建者：Done-0
// 创建时间：2025-08-13
package consts

// 文章状态常量
const (
	PostStatusDraft     = "draft"     // 草稿状态 - 文章正在编辑中，不对外展示
	PostStatusPublished = "published" // 已发布状态 - 文章已发布，对外可见
	PostStatusPrivate   = "private"   // 私有状态 - 文章仅作者可见
	PostStatusArchived  = "archived"  // 已归档状态 - 文章已归档，不在列表中显示但可通过链接访问
)
