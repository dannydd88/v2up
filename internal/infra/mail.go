package infra

import (
	gomail "gopkg.in/mail.v2"
)

type Mailer struct {
	smtpAdress   string
	smtpPort     int
	smtpUsername string
	smtpPassword string
}

func (m *Mailer) SendMail(to, subject, msg string) error {
	mail := gomail.NewMessage()
	mail.SetHeader("From", m.smtpUsername)
	mail.SetHeader("To", to)
	mail.SetHeader("Subject", subject)
	mail.SetBody("text/plain", msg)
	d := gomail.NewDialer(m.smtpAdress, m.smtpPort, m.smtpUsername, m.smtpPassword)
	err := d.DialAndSend(mail)
	return err
}
