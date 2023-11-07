package email

import (
	"fmt"
	"net/smtp"

	"github.com/jordan-wright/email"
)

func SendEmail(smtpHost, smtpPort, password, from, subject, html_content string, to []string) error {
	e := email.NewEmail()
	e.From = from
	e.To = to
	e.Subject = subject
	e.HTML = []byte(html_content)
	err := e.Send(smtpHost+":"+smtpPort, smtp.PlainAuth("", from, password, smtpHost))
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
