package wxmail_test

import (
	"github.com/kingson4wu/mp_weixin_server/mail"
	"github.com/kingson4wu/mp_weixin_server/weixin/wxmail"
	"testing"
)

func TestSendMail(t *testing.T) {
	wxmail.SendMail("oqV-XjlEcZZcA4pCwoaiLtnFF0XQ", "服务重启", "xxx", []mail.Attachment{})
}
