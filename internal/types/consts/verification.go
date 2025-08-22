// Package consts 提供验证码相关常量定义
// 创建者：Done-0
// 创建时间：2025-08-10
package consts

import "time"

const (
	// 验证码缓存键前缀
	EmailVerificationKeyPrefix = "verification:email" // 邮箱验证码缓存键前缀: verification:email:{email}
	ImgVerificationKeyPrefix   = "verification:image" // 图形验证码缓存键前缀: verification:image:{sessionId}

	// 验证码缓存过期时间
	EmailVerificationExpiration = 3 * time.Minute // 邮箱验证码缓存过期时间（3分钟）
	ImgVerificationExpiration   = 5 * time.Minute // 图形验证码缓存过期时间（5分钟）

	// 验证码长度配置
	EmailVerificationLength = 6 // 邮箱验证码长度（6位数字）
	ImgVerificationLength   = 4 // 图形验证码长度（4位字符）
)
