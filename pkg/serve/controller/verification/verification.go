package verification

import (
	"context"
	"net/http"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
	bizerr "jank.com/jank_blog/internal/error"
	"jank.com/jank_blog/internal/global"
	"jank.com/jank_blog/internal/utils"
	"jank.com/jank_blog/pkg/vo"
)

const (
	EmailVerificationCodeCacheKeyPrefix  = "Email:VERIFICATION:CODE:"
	EmailVerificationCodeCacheExpiration = 180 * time.Second
	ImgVerificationCodeCachePrefix       = "IMG:VERIFICATION:CODE:CACHE:"
	ImgVerificationCodeCacheExpiration   = 180 * time.Second
)

var ctx = context.Background()

// GenImgVerificationCode godoc
// @Summary      生成图形验证码并返回Base64编码
// @Description  生成单个图形验证码并将其返回为Base64编码字符串，用户可以用该验证码进行校验。
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        email  query   string  true  "邮箱地址，用于生成验证码"
// @Success      200   {object} vo.Result{data=map[string]string} "成功返回验证码的Base64编码"
// @Failure      400   {object} vo.Result{data=string} "请求参数错误，邮箱地址为空"
// @Failure      500   {object} vo.Result{data=string} "服务器错误，生成验证码失败"
// @Router       /account/genImgVerificationCode [get]
func GenImgVerificationCode(c echo.Context) error {
	email := c.QueryParam("email")
	if email == "" {
		utils.BizLogger(c).Errorf("请求参数错误，邮箱地址为空")
		return c.JSON(http.StatusBadRequest, vo.Fail("请求参数错误，邮箱地址为空", bizerr.New(bizerr.UnKnowErr), c))
	}

	key := ImgVerificationCodeCachePrefix + email

	// 生成单个图形验证码
	imgBase64, answer, err := utils.GenImgVerificationCode(email)
	if err != nil {
		utils.BizLogger(c).Errorf("生成图片验证码失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail("服务器错误，生成验证码失败", bizerr.New(bizerr.ServerError), c))
	}

	// 将验证码存储在 Redis 中，并设置过期时间
	err = global.Redis.Set(ctx, key, answer, ImgVerificationCodeCacheExpiration).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("图形验证码写入缓存失败，key: %v, 错误: %v", key, err)
		return c.JSON(http.StatusInternalServerError, vo.Fail("服务器错误，生成图形验证码失败", bizerr.New(bizerr.ServerError), c))
	}

	return c.JSON(http.StatusOK, vo.Success(map[string]string{"imgBase64": imgBase64}, c))
}

// SendEmailVerificationCode godoc
// @Summary      发送邮箱验证码
// @Description  向指定邮箱发送验证码，验证码有效期为3分钟
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        email  query   string  true  "邮箱地址，用于发送验证码"
// @Success      200     {object} vo.Result  "验证码发送成功, 请注意查收邮件"
// @Failure      400     {object} vo.Result  "请求参数错误，邮箱地址为空"
// @Failure      500     {object} vo.Result  "服务器错误，验证码发送失败"
// @Router       /account/sendEmailVerificationCode [get]
func SendEmailVerificationCode(c echo.Context) error {
	req := c.QueryParam("email")
	if req == "" {
		utils.BizLogger(c).Errorf("请求参数错误，邮箱地址为空")
		return c.JSON(http.StatusBadRequest, vo.Fail("请求参数错误，邮箱地址为空", bizerr.New(bizerr.SendEmailVerificationCodeFail), c))
	}

	key := EmailVerificationCodeCacheKeyPrefix + req

	// 检查验证码是否存在并有效
	exists, err := global.Redis.Exists(ctx, key).Result()
	if err != nil {
		utils.BizLogger(c).Errorf("检查邮箱验证码是否有效失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(nil, bizerr.New(bizerr.ServerError), c))
	}

	if exists > 0 {
		return c.JSON(http.StatusBadRequest, vo.Fail(nil, bizerr.New(bizerr.SendEmailVerificationCodeFail), c))
	}

	// 生成新的验证码并保存到缓存
	code := utils.NewRand()
	err = global.Redis.Set(ctx, key, strconv.Itoa(code), EmailVerificationCodeCacheExpiration).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("邮箱验证码写入缓存失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(nil, bizerr.New(bizerr.ServerError), c))
	}

	// 发送邮件
	if utils.ValidEmail(req) {
		emailContent := "您的注册验证码是: " + strconv.Itoa(code) + " , 有效期为 3 分钟"
		utils.SendEmail(emailContent, []string{req})
	}

	return c.JSON(http.StatusOK, vo.Success("邮箱验证码发送成功, 请注意查收邮件", c))
}

// VerifyEmailCode 检查邮箱验证码
func VerifyEmailCode(code, email string, c echo.Context) bool {
	return verifyCode(code, email, EmailVerificationCodeCacheKeyPrefix, c)
}

// VerificationImgCode 校验图形验证码
func VerificationImgCode(code, email string, c echo.Context) bool {
	return verifyCode(code, email, ImgVerificationCodeCachePrefix, c)
}

// verifyCode 验证码校验函数
func verifyCode(code, email, prefix string, c echo.Context) bool {
	key := prefix + email

	storedCode, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			utils.BizLogger(c).Error("验证码不存在或已过期")
		} else {
			utils.BizLogger(c).Errorf("验证码校验失败: %v", err)
		}
		return false
	}

	if storedCode != code {
		utils.BizLogger(c).Error("验证码错误")
		return false
	}

	// 验证通过后删除缓存
	if err := global.Redis.Del(ctx, key).Err(); err != nil {
		utils.BizLogger(c).Errorf("删除验证码缓存失败: %v", err)
	}

	return true
}
