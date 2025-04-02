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

	// Adicionando cabeçalhos obrigatórios para evitar problemas de entrega
	msg := []byte(fmt.Sprintf(
		"From: %s\r\nTo: %s\r\nSubject: %s\r\nContent-Type: text/plain; charset=UTF-8\r\n\r\n%s",
		e.sender, to, subject, body,
	))

	err := smtp.SendMail(e.smtpHost+":"+e.smtpPort, auth, e.sender, []string{to}, msg)
	if err != nil {
		fmt.Printf("Error to send mail: %v", err)
		return err
	}

	return nil
}
