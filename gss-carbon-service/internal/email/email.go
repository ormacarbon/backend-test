package email

import (
	"go.uber.org/zap"
)

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct {
	logger *zap.SugaredLogger
}

func NewEmailService(logger *zap.SugaredLogger) EmailService {
	return &emailService{
		logger: logger.Named("EmailService"),
	}
}

func (s *emailService) SendEmail(to, subject, body string) error {
	s.logger.Infow("Simulating email send",
		"to", to,
		"subject", subject,
	)

	return nil
}
