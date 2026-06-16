package mail

import (
	"sync"

	"gonotiserv/config"

	"gopkg.in/gomail.v2"
)

var (
	sender gomail.SendCloser
	mu     sync.Mutex
	cfg    *config.ENV
)

func InitSMTP(c *config.ENV) error {

	cfg = c

	dialer := gomail.NewDialer(
		cfg.SMTPHost,
		587,
		cfg.SMTPEmail,
		cfg.SMTPPassword,
	)

	s, err := dialer.Dial()
	if err != nil {
		return err
	}

	sender = s

	return nil
}

func SendEmail(
	to string,
	subject string,
	body string,
) error {

	m := gomail.NewMessage()

	m.SetHeader(
		"From",
		cfg.SMTPEmail,
	)

	m.SetHeader(
		"To",
		to,
	)

	m.SetHeader(
		"Subject",
		subject,
	)

	m.SetBody(
		"text/html",
		body,
	)

	mu.Lock()
	defer mu.Unlock()

	return gomail.Send(
		sender,
		m,
	)
}

func CloseSMTP() error {

	if sender != nil {
		return sender.Close()
	}

	return nil
}
