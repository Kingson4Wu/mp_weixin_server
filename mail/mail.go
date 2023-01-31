package mail

import (
	"fmt"
	"mime"
	"path/filepath"

	"gopkg.in/gomail.v2"
)

type Sender struct {
	mailAddress string // 发送地址
	mailName    string // 发送名称
	mailPass    string
	smtpHost    string
}

func New(mailAddress, mailName, mailPass, smtpHost string) *Sender {
	return &Sender{
		mailAddress: mailAddress,
		mailName:    mailName,
		mailPass:    mailPass,
		smtpHost:    smtpHost,
	}
}

type Attachment struct {
	FilePath string
	Name     string
}

func (s *Sender) SendMail(mailTo []string, subject string, body string) error {
	return s.SendMailWithAttachment(mailTo, subject, body, []Attachment{})
}

func (s *Sender) SendMailWithAttachment(mailTo []string, subject string, body string, attachments []Attachment) error {

	m := gomail.NewMessage(
		gomail.SetEncoding(gomail.Base64),
	)
	if s.mailName != "" {
		m.SetHeader("From", m.FormatAddress(s.mailAddress, s.mailName))
	} else {
		m.SetHeader("From", s.mailAddress)
	}

	m.SetHeader("To", mailTo...)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)

	for _, attachment := range attachments {

		fileName := attachment.Name
		if fileName == "" {
			fileName = filepath.Base(attachment.FilePath)
		}
		m.Attach(attachment.FilePath,
			gomail.Rename(fileName),
			gomail.SetHeader(map[string][]string{
				"Content-Disposition": {
					fmt.Sprintf(`attachment; filename="%s"`, mime.QEncoding.Encode("UTF-8", fileName)),
				},
			}),
		)

	}

	/*
	   创建SMTP客户端，连接到远程的邮件服务器，需要指定服务器地址、端口号、用户名、密码，如果端口号为465的话，
	   自动开启SSL，这个时候需要指定TLSConfig
	*/
	d := gomail.NewDialer(s.smtpHost, 465, s.mailAddress, s.mailPass)
	//d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	err := d.DialAndSend(m)
	return err
}
