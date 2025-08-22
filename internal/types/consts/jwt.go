// Package consts 提供JWT相关常量定义
// 创建者：Done-0
// 创建时间：2025-08-12
package consts

const (
	// JWT 相关常量
	JWTRealm        = "jank" // JWT 领域标识符
	JWTSubjectClaim = "sub"  // JWT 标准主体声明键: 存储用户 ID

	// JWT 认证相关常量
	JWTTokenLookup   = "header:Authorization" // JWT token 查找位置
	JWTTokenHeadName = "Bearer"               // JWT token 头部名称
	JWTBearerPrefix  = "Bearer "              // JWT Bearer 前缀（含空格）
)
