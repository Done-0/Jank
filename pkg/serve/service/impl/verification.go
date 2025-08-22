// Package impl 提供验证码相关的业务逻辑实现
// 创建者：Done-0
// 创建时间：2025-08-10
package impl

import (
	"context"
	"fmt"
	"strconv"

	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/internal/global"
	"github.com/Done-0/jank/internal/types/consts"
	"github.com/Done-0/jank/internal/utils/email"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
)

// VerificationServiceImpl 验证码服务实现
type VerificationServiceImpl struct{}

// NewVerificationService 创建验证码服务实例
func NewVerificationService() service.VerificationService {
	return &VerificationServiceImpl{}
}

// SendEmailVerificationCode 发送邮箱验证码
func (v *VerificationServiceImpl) SendEmailVerificationCode(c *app.RequestContext, req *dto.SendEmailCodeRequest) error {
	key := consts.EmailVerificationKeyPrefix + req.Email

	// 生成并缓存验证码（如果已存在则覆盖旧的验证码）
	code := email.NewRand()
	err := global.RedisClient.Set(context.Background(), key, strconv.Itoa(code), consts.EmailVerificationExpiration).Err()
	if err != nil {
		return err
	}

	// 发送验证码邮件
	expirationInMinutes := int(consts.EmailVerificationExpiration.Minutes())
	emailContent := fmt.Sprintf("Your registration verification code is: %d, valid for %d minutes.", code, expirationInMinutes)
	success, err := email.SendEmail(emailContent, []string{req.Email})
	if !success {
		global.RedisClient.Del(context.Background(), key)
		return err
	}

	return nil
}
