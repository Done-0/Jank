// Package service 提供验证码相关的服务层接口
// 创建者：Done-0
// 创建时间：2025-08-10
package service

import (
	"github.com/cloudwego/hertz/pkg/app"

	"github.com/Done-0/jank/pkg/serve/controller/dto"
)

// VerificationService 验证码服务接口
type VerificationService interface {
	SendEmailVerificationCode(c *app.RequestContext, req *dto.SendEmailCodeRequest) error // 发送邮箱验证码
}
