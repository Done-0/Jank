// Package consts 提供缓存相关常量定义
// 创建者：Done-0
// 创建时间：2025-08-10
package consts

const (
	// Redis 缓存键前缀 - 认证相关
	AuthAccessTokenKeyPrefix  = "auth:access_token"  // 访问令牌缓存键前缀: auth:access_token:{userID}
	AuthRefreshTokenKeyPrefix = "auth:refresh_token" // 刷新令牌缓存键前缀: auth:refresh_token:{userID}
)
