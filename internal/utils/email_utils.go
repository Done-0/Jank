// Package utils 提供邮件操作相关工具
// 创建者：Done-0
// 创建时间：2025-05-10
package utils

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"regexp"
	"time"

	"github.com/jordan-wright/email"

	"jank.com/jank_blog/configs"
	"jank.com/jank_blog/internal/global"
)

const SUBJECT = "【Jank Blog】注册验证码"

// 邮箱服务器配置
var emailServers = map[string]struct {
	Server, Port string
}{
	"qq":      {"smtp.qq.com", ":587"},
	"gmail":   {"smtp.gmail.com", ":587"},
	"outlook": {"smtp.office365.com", ":587"},
}

// SendEmail 发送邮件到指定邮箱
// 参数：
//   - content: 邮件内容
//   - toEmail: 接收邮件的邮箱地址数组
//
// 返回值：
//   - bool: 发送成功返回 true，失败返回 false
//   - error: 发送过程中的错误
func SendEmail(content string, toEmail []string) (bool, error) {
	config, err := configs.LoadConfig()
	if err != nil {
		global.SysLog.Errorf("加载邮件配置失败, toEmail: %v, 错误信息: %v", toEmail, err)
		return false, fmt.Errorf("加载邮件配置失败: %v", err)
	}

	// 获取SMTP相关配置
	fromEmail := config.AppConfig.FromEmail
	emailType := config.AppConfig.EmailType

	// 获取邮箱类型和对应的服务器配置
	serverConfig, ok := emailServers[emailType]
	if !ok || emailType == "" {
		emailType = "qq"
		serverConfig = emailServers[emailType]
		global.SysLog.Warnf("邮箱类型无效或为空, 原类型: %s, 默认使用 QQ 邮箱替代", emailType)
	}

	e := email.NewEmail()
	e.From = fromEmail
	e.To = toEmail
	e.Subject = SUBJECT
	e.Text = []byte(content)

	smtpAddr := serverConfig.Server + serverConfig.Port
	auth := smtp.PlainAuth("", fromEmail, config.AppConfig.EmailSmtp, serverConfig.Server)

	if err := e.Send(smtpAddr, auth); err != nil {
		global.SysLog.Errorf("发送邮件失败, toEmail: %v, 错误信息: %v", toEmail, err)
		return false, fmt.Errorf("发送邮件失败: %v", err)
	}

	return true, nil
}

// NewRand 生成六位数随机验证码
// 返回值：
//   - int: 六位数随机验证码
func NewRand() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}

// ValidEmail 检查邮箱格式是否有效
// 参数：
//   - email: 待验证的邮箱地址
//
// 返回值：
//   - bool: 邮箱格式有效返回 true，无效返回 false
func ValidEmail(email string) bool {
	pattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	return regexp.MustCompile(pattern).MatchString(email)
}
