package wxmail

import (
	"fmt"
	"github.com/kingson4wu/mp_weixin_server/global"
	"github.com/kingson4wu/mp_weixin_server/mail"
)

func SendMail(account string, subject string, body string, attachments []mail.Attachment) {

	if v, ok := global.OpenidToMail[account]; ok {

		mailTo := []string{v}

		if len(attachments) > 0 {
			err := global.MailSender.SendMailWithAttachment(mailTo, subject, body, attachments)
			if err != nil {
				fmt.Println("Send fail! - ", err)
				return
			}
		} else {
			err := global.MailSender.SendMail(mailTo, subject, body)
			if err != nil {
				fmt.Println("Send fail! - ", err)
				return
			}
		}
	}
}
