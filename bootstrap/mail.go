package bootstrap

import (
	"context"
	"crypto/tls"
	"os"
	"strconv"

	"github.com/Decem-Technology/service-helper/proto/mail"
	"gopkg.in/gomail.v2"
)

type (
	// Mailer mailer
	Mailer struct {
	}
)

var mailer *gomail.Dialer

// CreateMailerConnection create mailer connection
func CreateMailerConnection() {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic("[SMTP] smtp port is invalid")
	}
	d := gomail.NewDialer(
		os.Getenv("MAIL_SMTP"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)
	if os.Getenv("MAIL_INSECURE") == "true" {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	} else if os.Getenv("MAIL_INSECURE") == "false" {
		d.TLSConfig = &tls.Config{InsecureSkipVerify: false}
	}
	mailer = d
}

// Mail get mailer fn
func (ctl *Mailer) Mail() *gomail.Dialer {
	return mailer
}

// Send send mail
func (ctl *Mailer) Send(ct context.Context, mail *mail.MailMessage) error {
	m := gomail.NewMessage()
	m.SetAddressHeader("From", mail.From, mail.Sender)
	if len(mail.To) == 1 {
		m.SetAddressHeader("To", mail.To[0], mail.Receiver)
	} else {
		m.SetHeader("To", mail.To...)
	}
	if len(mail.Cc) > 0 {
		m.SetHeader("Cc", mail.Cc...)
	}
	if len(mail.Bcc) > 0 {
		m.SetHeader("Bcc", mail.Bcc...)
	}
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)
	if len(mail.Attach) > 0 {
		for _, i := range mail.Attach {
			m.Attach(i)
		}
	}
	if err := ctl.Mail().DialAndSend(m); err != nil {
		return err
	}
	return nil
}
