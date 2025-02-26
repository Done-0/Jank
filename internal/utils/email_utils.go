package utils

import (
	"math/rand"
	"net/smtp"
	"regexp"
	"time"

	"github.com/jordan-wright/email"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

const (
	SUBJECT         = "【Jank Blog】注册验证码"
	QqEmailSmtp     = "smtp.qq.com"
	QqEmailSmtpPort = ":25"
)

func SendEmail(content string, toEmail []string) bool {
	if content == "" || len(toEmail) == 0 {
		return false
	}

	config, err := configs.LoadConfig()
	if err != nil {
		global.SysLog.Errorf("加载 SMTP Auth 配置失败, toEmail: %v, 错误信息: %v", toEmail, err)
		return false
	}

	ema := email.NewEmail()
	ema.From = config.AppConfig.FromEmail
	ema.To = toEmail
	ema.Subject = SUBJECT
	ema.Text = []byte(content)
	err = ema.Send(QqEmailSmtp+QqEmailSmtpPort, smtp.PlainAuth("", config.AppConfig.FromEmail, config.AppConfig.QqSmtp, QqEmailSmtp))
	if err != nil {
		global.SysLog.Errorf("发送邮件失败, toEmail: %v, 错误信息: %v", toEmail, err)
		return false
	}

	return true
}

// NewRand 生成一个六位数的随机验证码
func NewRand() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}

// ValidEmail 检查邮箱格式是否有效
func ValidEmail(email string) bool {
	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	return re.MatchString(email)
}
