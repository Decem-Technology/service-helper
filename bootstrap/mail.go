package bootstrap

import (
	"fmt"
	"os"
	"strconv"

	"gopkg.in/gomail.v2"
)

type (
	// Mailer mailer
	Mailer struct {
	}
)

var mailer *gomail.Dialer

// CreateMailerConnection make mailer connection
func CreateMailerConnection() {
	port, err := strconv.Atoi(os.Getenv("MAIL_PORT"))
	if err != nil {
		panic("[SMTP] smtp port is invalid")
	}
	d := gomail.NewPlainDialer(
		os.Getenv("MAIL_SMTP"),
		port,
		os.Getenv("MAIL_USERNAME"),
		os.Getenv("MAIL_PASSWORD"),
	)
	// d.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	mailer = d
	fmt.Println("[Mailer] connected")
}

// Mail get mail
func (ctl *Mailer) Mail() *gomail.Dialer {
	return mailer
}
