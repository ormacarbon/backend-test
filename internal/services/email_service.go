package services

import (
	"fmt"
	"net/smtp"
	"os"
)

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}

type emailService struct {
	smtpHost string
	smtpPort string
	sender   string
	password string
}

func NewEmailService() EmailService {
	return &emailService{
		smtpHost: os.Getenv("SMTP_HOST"),
		smtpPort: os.Getenv("SMTP_PORT"),
		sender:   os.Getenv("SMTP_EMAIL"),
		password: os.Getenv("SMTP_PASSWORD"),
	}
}

func (e *emailService) SendEmail(to string, subject string, body string) error {
	auth := smtp.PlainAuth("", e.sender, e.password, e.smtpHost)

	msg := []byte(fmt.Sprintf("Subject: %s\n\n%s", subject, body))

	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.sender, []string{to}, msg)
	return err
}

