package global

import (
	"github.com/kingson4wu/mp_weixin_server/config"
	"github.com/kingson4wu/mp_weixin_server/mail"
)

var (
	MailSender   *mail.Sender
	OpenidToMail map[string]string
)

func init() {
	mailConfig := config.GetMailConfig()
	MailSender = mail.New(mailConfig.MailAddress, mailConfig.MailName, mailConfig.MailPass, mailConfig.SmtpHost)

	elements := mailConfig.UserMailInfos
	OpenidToMail = make(map[string]string)
	for _, data := range elements {
		OpenidToMail[data.OpenId] = data.Address
	}
}
