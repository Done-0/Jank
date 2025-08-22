// Package controller 提供验证码相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-08-10
package controller

import (
	"context"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"

	"github.com/Done-0/jank/internal/types/errno"
	"github.com/Done-0/jank/internal/utils/errorx"
	"github.com/Done-0/jank/internal/utils/validator"
	"github.com/Done-0/jank/internal/utils/vo"
	"github.com/Done-0/jank/pkg/serve/controller/dto"
	"github.com/Done-0/jank/pkg/serve/service"
)

// VerificationController 验证码控制器
type VerificationController struct {
	verificationService service.VerificationService
}

// NewVerificationController 创建验证码控制器
// 参数：
//
//	verificationService: 验证码服务
//
// 返回值：
//
//	*VerificationController: 验证码控制器
func NewVerificationController(verificationService service.VerificationService) *VerificationController {
	return &VerificationController{
		verificationService: verificationService,
	}
}

// SendEmailVerificationCode 发送邮箱验证码
// @Router /api/v1/verification/email [get]
func (ctrl *VerificationController) SendEmailVerificationCode(ctx context.Context, c *app.RequestContext) {
	req := new(dto.SendEmailCodeRequest)
	if err := c.BindQuery(req); err != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, err, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "bind query failed"))))
		return
	}

	errors := validator.Validate(req)
	if errors != nil {
		c.JSON(consts.StatusBadRequest, vo.Fail(c, errors, errorx.New(errno.ErrInvalidParams, errorx.KV("msg", "validation failed"))))
		return
	}

	err := ctrl.verificationService.SendEmailVerificationCode(c, req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, vo.Fail(c, err, errorx.New(errno.ErrInternalServer, errorx.KV("msg", "send email verification code failed"))))
		return
	}

	c.JSON(consts.StatusOK, vo.Success(c, "verification code sent successfully, please check your email"))
}
