package global

import (
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/mail"
)

var MailSender *mail.Sender

func init() {
	mailConfig := config.GetMailConfig()
	MailSender = mail.New(mailConfig.MailAddress, mailConfig.MailName, mailConfig.MailPass, mailConfig.SmtpHost)
}
