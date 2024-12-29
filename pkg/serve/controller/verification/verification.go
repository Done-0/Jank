package account

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
	VerificationCodeCacheKeyPrefix     = "Email:VERIFICATION:CODE:"
	VerificationCodeCacheExpiration    = 180 * time.Second
	ImgVerificationCodeCachePrefix     = "IMG:VERIFICATION:CODE:CACHE:"
	ImgVerificationCodeCacheExpiration = 180 * time.Second
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
	registerParam := c.QueryParam("email")
	if registerParam == "" {
		utils.BizLogger(c).Errorf("请求参数错误，邮箱地址为空")
		return c.JSON(http.StatusBadRequest, vo.Fail("请求参数错误，邮箱地址为空", bizerr.New(bizerr.UnKnowErr), c))
	}

	key := ImgVerificationCodeCachePrefix + registerParam

	// 生成单个图形验证码
	imgBase64, err := utils.GenImgVerificationCode(registerParam)
	if err != nil {
		utils.BizLogger(c).Errorf("生成图片验证码失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail("服务器错误，生成验证码失败", bizerr.New(bizerr.ServerError), c))
	}

	// 将验证码存储在 Redis 中，并设置过期时间
	err = global.Redis.Set(ctx, key, imgBase64, ImgVerificationCodeCacheExpiration).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("验证码写入缓存失败，key: %v, 错误: %v", key, err)
		return c.JSON(http.StatusInternalServerError, vo.Fail("服务器错误，生成验证码失败", bizerr.New(bizerr.ServerError), c))
	}

	utils.BizLogger(c).Info("图形验证码生成成功")
	return c.JSON(http.StatusOK, vo.Success(map[string]string{"imgBase64": imgBase64}, c))
}

// SendEmailVerificationCode godoc
// @Summary      发送邮箱验证码
// @Description  向指定邮箱发送验证码，验证码有效期为3分钟
// @Tags         账户
// @Accept       json
// @Produce      json
// @Param        email  query   string  true  "邮箱地址，用于发送验证码"
// @Success      200     {object} vo.Result  "成功发送验证码"
// @Failure      400     {object} vo.Result  "请求参数错误，邮箱地址为空或验证码仍然有效"
// @Failure      500     {object} vo.Result  "服务器错误，验证码发送失败"
// @Router       /account/sendEmailVerificationCode [get]
func SendEmailVerificationCode(c echo.Context) error {
	target := c.QueryParam("email")
	if target == "" {
		utils.BizLogger(c).Errorf("验证码目标地址为空")
		return c.JSON(http.StatusBadRequest, vo.Fail(nil, bizerr.New(bizerr.SendEmailVerificationCodeFail), c))
	}

	key := VerificationCodeCacheKeyPrefix + target

	// 检查验证码是否存在并有效
	exists, err := global.Redis.Exists(ctx, key).Result()
	if err != nil {
		utils.BizLogger(c).Errorf("检查验证码是否有效失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(nil, bizerr.New(bizerr.ServerError), c))
	}

	if exists > 0 {
		return c.JSON(http.StatusBadRequest, vo.Fail(nil, bizerr.New(bizerr.SendEmailVerificationCodeFail), c))
	}

	// 生成新的验证码并保存到缓存
	code := utils.NewRand()
	err = global.Redis.Set(ctx, key, strconv.Itoa(code), VerificationCodeCacheExpiration).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("验证码写入缓存失败: %v", err)
		return c.JSON(http.StatusInternalServerError, vo.Fail(nil, bizerr.New(bizerr.ServerError), c))
	}

	// 发送邮件
	if utils.ValidEmail(target) {
		emailContent := "您的注册验证码是: " + strconv.Itoa(code) + " , 有效期为 3 分钟"
		utils.SendEmail(emailContent, []string{target})
	}

	return c.JSON(http.StatusOK, vo.Success("验证码发送成功, 请注意查收邮件", c))
}

// VerifyEmailCode 检查提供的验证码是否与存储的验证码匹配
func VerifyEmailCode(code, email string, c echo.Context) bool {
	key := VerificationCodeCacheKeyPrefix + email

	// 检查验证码缓存是否存在
	exists, err := global.Redis.Exists(ctx, key).Result()
	if err != nil {
		utils.BizLogger(c).Errorf("检查验证码缓存失败: %v", err)
		return false
	}

	if exists == 0 {
		utils.BizLogger(c).Error("验证码不存在或已过期")
		return false
	}

	// 获取缓存中的验证码
	storedCode, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		utils.BizLogger(c).Errorf("获取验证码缓存失败: %v", err)
		return false
	}

	// 比对验证码
	if storedCode != code {
		utils.BizLogger(c).Error("验证码不匹配")
		return false
	}

	// 删除验证码缓存
	err = global.Redis.Del(ctx, key).Err()
	if err != nil {
		utils.BizLogger(c).Errorf("删除验证码缓存失败: %v", err)
		return false
	}

	return true
}

// VerificationImgCode 校验图形验证码
func VerificationImgCode(code, email string, c echo.Context) bool {
	key := ImgVerificationCodeCachePrefix + email

	// 从Redis获取存储的验证码
	storedCode, err := global.Redis.Get(ctx, key).Result()
	if err != nil {
		if err.Error() == "redis: nil" {
			utils.BizLogger(c).Error("图形验证码不存在或已过期")
		} else {
			utils.BizLogger(c).Errorf("获取图形验证码失败: %v", err)
		}
		return false
	}

	isValid := storedCode == code
	if !isValid {
		utils.BizLogger(c).Error("图形验证码不匹配")
	} else {
		if err := global.Redis.Del(ctx, key).Err(); err != nil {
			utils.BizLogger(c).Errorf("删除已使用的图形验证码失败: %v", err)
		}
	}

	return isValid
}
