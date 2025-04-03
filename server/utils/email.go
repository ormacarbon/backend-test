package utils

import (
	"fmt"
	"net/smtp"
	"os"
)

func checkSMTPEnv() error {
	requiredVars := []string{
		"SMTP_HOST",
		"SMTP_PORT",
		"SMTP_USER",
		"SMTP_PASSWORD",
	}

	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			return fmt.Errorf("required SMTP environment variable %s is not set", v)
		}
	}
	return nil
}

func SendEmail(to, subject, body string) error {
	if err := checkSMTPEnv(); err != nil {
		return fmt.Errorf("email configuration error: %w", err)
	}

	from := os.Getenv("SMTP_USER")
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := os.Getenv("SMTP_HOST")
	smtpPort := os.Getenv("SMTP_PORT")

	msg := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s",
		from, to, subject, body)

	auth := smtp.PlainAuth("", from, password, smtpHost)

	err := smtp.SendMail(smtpHost+":"+smtpPort,
		auth,
		from,
		[]string{to},
		[]byte(msg))

	if err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}
