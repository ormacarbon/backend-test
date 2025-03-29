package services

import (
	"fmt"
	"log"

	"github.com/joaooliveira247/backend-test/src/config"
	"gopkg.in/gomail.v2"
)

func SendEmail(to, subject, message string) {
	smtpHost := "smtp.gmail.com"
	smtpPort := 587

	msg := gomail.NewMessage()

	headers := map[string]string{
		"From":    config.SERVICE_EMAIL,
		"To":      to,
		"Subject": subject,
	}

	for header, value := range headers {
		msg.SetHeader(header, value)
	}

	msg.SetBody(
		"text/html",
		fmt.Sprintf("<h1>Carbon Offset Competition</h1><p>%s</p>", message),
	)

	server := gomail.NewDialer(
		smtpHost,
		smtpPort,
		config.SERVICE_EMAIL,
		config.PASSWORD_SERVICE_EMAIL,
	)

	if err := server.DialAndSend(msg); err != nil {
		log.Fatalf("Send Email: %v", err)
	}
}
