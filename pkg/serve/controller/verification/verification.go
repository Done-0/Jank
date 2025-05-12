// Package verification 提供验证码相关的HTTP接口处理
// 创建者：Done-0
// 创建时间：2025-05-10
package verification

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/labstack/echo/v4"

	bizErr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/vo"
	"jank.com/jank_blog/pkg/vo/verification"
)

const (
	EMAIL_VERIFICATION_CODE_CACHE_KEY_PREFIX = "EMAIL:VERIFICATION:CODE:"     // 邮箱验证码缓存前缀
	EMAIL_VERIFICATION_CODE_CACHE_EXPIRATION = 3 * time.Minute                // 邮箱验证码缓存过期时间
	IMG_VERIFICATION_CODE_CACHE_PREFIX       = "IMG:VERIFICATION:CODE:CACHE:" // 图形验证码缓存前缀
	IMG_VERIFICATION_CODE_CACHE_EXPIRATION   = 3 * time.Minute                // 图形验证码缓存过期时间
)

// SendImgVerificationCode godoc
// @Summary      生成图形验证码并返回Base64编码
// @Description  生成单个图形验证码并将其返回为Base64编码字符串，用户可以用该验证码进行校验。
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        email  query   string  true  "邮箱地址，用于生成验证码"
// @Success      200   {object} vo.Result{data=map[string]string} "成功返回验证码的Base64编码"
// @Failure      400   {object} vo.Result{data=string} "请求参数错误，邮箱地址为空"
// @Failure      500   {object} vo.Result{data=string} "服务器错误，生成验证码失败"
// @Router       /verification/sendImgVerificationCode [get]
func SendImgVerificationCode(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		utils.BizLogger(c).Errorf("请求参数错误，邮箱地址为空")
		return c.JSON(http.StatusBadRequest, vo.Fail(c, "请求参数错误，邮箱地址为空", bizErr.New(bizErr.BAD_REQUEST)))
	}

	key := IMG_VERIFICATION_CODE_CACHE_PREFIX + email

	// 生成单个图形验证码
	imgBase64, answer, err := utils.GenImgVerificationCode()
	if err != nil {
		utils.BizLogger(c).Errorf("生成图片验证码失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR)))
	}

	err = global.RedisClient.Set(context.Background(), key, answer, IMG_VERIFICATION_CODE_CACHE_EXPIRATION).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("图形验证码写入缓存失败，key: %v, 错误: %v", key, err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR)))
	}

	return c.JSON(http.StatusOK, vo.Success(c, verification.ImgVerificationVO{ImgBase64: imgBase64}))
}

// SendEmailVerificationCode godoc
// @Summary 发送邮箱验证码
// @Description 向指定邮箱发送验证码，验证码有效期为3分钟
// @Tags 账户
// @Accept json
// @Produce json
// @Param email query string true "邮箱地址，用于发送验证码"
// @Success 200 {object} vo.Result "邮箱验证码发送成功, 请注意查收邮件"
// @Failure 400 {object} vo.Result "请求参数错误，邮箱地址为空"
// @Failure 500 {object} vo.Result "服务器错误，邮箱验证码发送失败"
// @Router /verification/sendEmailVerificationCode [get]
func SendEmailVerificationCode(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		utils.BizLogger(c).Errorf("请求参数错误，邮箱地址为空")
		return c.JSON(http.StatusBadRequest, vo.Fail(c, "请求参数错误，邮箱地址为空", bizErr.New(bizErr.BAD_REQUEST)))
	}

	if !utils.ValidEmail(email) {
		utils.BizLogger(c).Errorf("邮箱格式无效: %s", email)
		return c.JSON(http.StatusBadRequest, vo.Fail(c, "邮箱格式无效", bizErr.New(bizErr.BAD_REQUEST)))
	}

	key := EMAIL_VERIFICATION_CODE_CACHE_KEY_PREFIX + email

	// 检查验证码是否存在
	exists, err := global.RedisClient.Exists(context.Background(), key).Result()
	if err != nil {
		utils.BizLogger(c).Errorf("检查邮箱验证码是否有效失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR)))
	}
	if exists > 0 {
		return c.JSON(http.StatusBadRequest, vo.Fail(c, "邮箱验证码已存在", bizErr.New(bizErr.SERVER_ERR)))
	}

	// 生成并缓存验证码
	code := utils.NewRand()
	err = global.RedisClient.Set(context.Background(), key, strconv.Itoa(code), EMAIL_VERIFICATION_CODE_CACHE_EXPIRATION).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("邮箱验证码写入缓存失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SERVER_ERR)))
	}

	// 发送验证码邮件
	expirationInMinutes := int(EMAIL_VERIFICATION_CODE_CACHE_EXPIRATION.Round(time.Minute).Minutes())
	emailContent := fmt.Sprintf("您的注册验证码是: %d , 有效期为 %d 分钟。", code, expirationInMinutes)
	success, err := utils.SendEmail(emailContent, []string{email})
	if !success {
		utils.BizLogger(c).Errorf("邮箱验证码发送失败，邮箱地址: %s, 错误: %v", email, err)
		global.RedisClient.Del(context.Background(), key)
		return c.JSON(http.StatusInternalServerError, vo.Fail(c, err, bizErr.New(bizErr.SEND_EMAIL_VERIFICATION_CODE_FAIL)))
	}

	return c.JSON(http.StatusOK, vo.Success(c, "邮箱验证码发送成功, 请注意查收！"))
}

// VerifyEmailCode 校验邮箱验证码
// 参数：
//   - c: Echo 上下文
//   - code: 验证码
//   - email: 邮箱地址
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyEmailCode(c echo.Context, code, email string) bool {
	return verifyCode(c, code, email, EMAIL_VERIFICATION_CODE_CACHE_KEY_PREFIX)
}

// VerifyImgCode 校验图形验证码
// 参数：
//   - c: Echo 上下文
//   - code: 验证码
//   - email: 邮箱地址
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func VerifyImgCode(c echo.Context, code, email string) bool {
	return verifyCode(c, code, email, IMG_VERIFICATION_CODE_CACHE_PREFIX)
}

// verifyCode 通用验证码校验
// 参数：
//   - c: Echo 上下文
//   - code: 验证码
//   - email: 邮箱地址
//   - prefix: 缓存键前缀
//
// 返回值：
//   - bool: 验证成功返回 true，失败返回 false
func verifyCode(c echo.Context, code, email, prefix string) bool {
	key := prefix + email

	storedCode, err := global.RedisClient.Get(c.Request().Context(), key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			utils.BizLogger(c).Error("验证码不存在或已过期")
		} else {
			utils.BizLogger(c).Errorf("验证码校验失败: %v", err)
		}
		return false
	}

	storedCode = strings.ToUpper(strings.TrimSpace(storedCode))
	code = strings.ToUpper(strings.TrimSpace(code))

	if storedCode != code {
		utils.BizLogger(c).Error("用户验证码错误")
		return false
	}

	if err := global.RedisClient.Del(context.Background(), key).Err(); err != nil {
		utils.BizLogger(c).Errorf("删除验证码缓存失败: %v", err)
	}

	return true
}
