package email

import "log"

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct{}

func NewEmailService() EmailService {
	return &emailService{}
}

func (s *emailService) SendEmail(to, subject, body string) error {
	// Exemplo: Apenas loga o envio do e-mail
	log.Printf("Sending email to %s: subject=%s, body=%s", to, subject, body)
	return nil
}
