package utils

import (
	"github.com/mojocn/base64Captcha"
	"jank.com/jank_blog/internal/global"
)

const (
	CaptchaSource   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789" // 验证码字符源
	IdPrefix        = "cap:jank_blog:"                                                 // 验证码存储键前缀
	FontFile        = "wqy-microhei.ttc"                                               // 字体文件
	ClearAfterCheck = true                                                             // 验证后是否清除验证码
	ImgHeight       = 80                                                               // 验证码图片高度
	ImgWidth        = 200                                                              // 验证码图片宽度
	NoiseCount      = 3                                                                // 干扰点数量
	CaptchaLength   = 4                                                                // 验证码字符长度
)

var store = base64Captcha.DefaultMemStore

// 生成图形验证码
func GenImgVerificationCode(verificationId string) (string, error) {
	driver := createDriver()
	captcha := base64Captcha.NewCaptcha(driver, store)

	_, content, answer := captcha.Driver.GenerateIdQuestionAnswer()
	if len(answer) > CaptchaLength {
		answer = answer[:CaptchaLength]
	}

	if err := captcha.Store.Set(IdPrefix+verificationId, answer); err != nil {
		global.SysLog.Errorf("生成图形验证码错误: %v", err)
		return "", err
	}

	item, _ := captcha.Driver.DrawCaptcha(content)
	return item.EncodeB64string(), nil
}

// 验证用户输入的验证码是否正确
func VerificationImgCode(verificationId string, code string) bool {
	return store.Verify(IdPrefix+verificationId, code, ClearAfterCheck)
}

// 创建验证码的驱动配置
func createDriver() *base64Captcha.DriverString {
	return &base64Captcha.DriverString{
		Height:          ImgHeight,
		Width:           ImgWidth,
		NoiseCount:      NoiseCount,
		ShowLineOptions: base64Captcha.OptionShowHollowLine | base64Captcha.OptionShowSineLine,
		Length:          CaptchaLength,
		Source:          CaptchaSource,
		Fonts:           []string{FontFile},
	}
}
