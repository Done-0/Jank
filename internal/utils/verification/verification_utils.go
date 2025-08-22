// Package verification 提供验证码校验工具函数
// 创建者：Done-0
// 创建时间：2025-08-10
package verification

import (
	"context"
	"strings"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/types/consts"
)

// VerifyEmailCode 校验邮箱验证码
// 参数：
//   - c: Hertz 请求上下文
//   - code: 验证码
//   - email: 邮箱地址
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyEmailCode(c *app.RequestContext, code, email string) bool {
	return VerifyCode(c, code, email, consts.EmailVerificationKeyPrefix)
}

// VerifyImgCode 校验图形验证码
// 参数：
//   - c: Hertz 请求上下文
//   - code: 验证码
//   - email: 邮箱地址
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyImgCode(c *app.RequestContext, code, email string) bool {
	return VerifyCode(c, code, email, consts.ImgVerificationKeyPrefix)
}

// VerifyCode 通用验证码校验
// 参数：
//   - c: Hertz 请求上下文
//   - code: 验证码
//   - email: 邮箱地址
//   - prefix: 缓存键前缀
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyCode(c *app.RequestContext, code, email, prefix string) bool {
	key := prefix + email

	storedCode, err := global.RedisClient.Get(context.Background(), key).Result()
	if err != nil {
		return false
	}

	storedCode = strings.ToUpper(strings.TrimSpace(storedCode))
	code = strings.ToUpper(strings.TrimSpace(code))

	if storedCode != code {
		return false
	}

	if err := global.RedisClient.Del(context.Background(), key).Err(); err != nil {
		return false
	}

	return true
}
