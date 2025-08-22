// Package email 提供邮件操作相关工具
// 创建者：Done-0
// 创建时间：2025-08-05
package email

import (
	"crypto/tls"
	"fmt"
	"math/rand"
	"time"

	"gopkg.in/gomail.v2"

	"github.com/Done-0/jank/configs"
	"github.com/Done-0/jank/internal/global"
)

// 邮件相关常量
const (
	EMAIL_SUBJECT = "【Jank Blog】注册验证码"
)

// 邮箱服务器配置
var emailServers = map[string]struct {
	Server string
	Port   int
	SSL    bool
}{
	"qq":      {"smtp.qq.com", 465, true},         // QQ 邮箱使用 SSL 加密
	"gmail":   {"smtp.gmail.com", 465, true},      // Gmail 使用 SSL 加密
	"outlook": {"smtp.office365.com", 587, false}, // Outlook 使用 TLS 加密
}

// SendEmail 发送邮件到指定邮箱
// 参数：
//
//	content: 邮件内容
//	toEmails: 目标邮箱
//
// 返回值：
//
//	bool: 发送成功返回 true，失败返回 false
//	error: 执行过程中的错误
func SendEmail(content string, toEmails []string) (bool, error) {
	cfgs, err := configs.GetConfig()
	if err != nil {
		global.SysLog.Errorf("failed to load email config: %v", err)
		return false, fmt.Errorf("failed to load email config: %v", err)
	}

	// 获取SMTP配置
	emailType := cfgs.AppConfig.Email.EmailType
	serverConfig := emailServers[emailType]

	// 创建邮件
	m := gomail.NewMessage()
	m.SetHeader("From", cfgs.AppConfig.Email.FromEmail)
	m.SetHeader("To", toEmails...)
	m.SetHeader("Subject", EMAIL_SUBJECT)
	m.SetBody("text/plain", content)

	// 配置发送器
	d := gomail.NewDialer(
		serverConfig.Server,
		serverConfig.Port,
		cfgs.AppConfig.Email.FromEmail,
		cfgs.AppConfig.Email.EmailSmtp,
	)

	// 根据端口配置安全选项
	if serverConfig.SSL {
		d.SSL = true
	} else {
		d.TLSConfig = &tls.Config{
			ServerName: serverConfig.Server,
			MinVersion: tls.VersionTLS12,
		}
	}

	if err := d.DialAndSend(m); err != nil {
		global.SysLog.Errorf("failed to send email: %v", err)
		return false, fmt.Errorf("failed to send email: %v", err)
	}

	return true, nil
}

// NewRand 生成六位数随机验证码
// 返回值：
//
//	int: 六位数随机数
func NewRand() int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(900000) + 100000
}
