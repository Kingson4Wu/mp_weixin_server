package mail_test

import (
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/mail"
	"testing"
)

func TestSendMail(t *testing.T) {
	mailConfig := config.GetMailConfig()
	sender := mail.New(mailConfig.MailAddress, mailConfig.MailName, mailConfig.MailPass, mailConfig.SmtpHost)
	sender.SendMail([]string{"819966354@qq.com"}, "来了！", "888")
}
