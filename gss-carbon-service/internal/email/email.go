package email

import (
	"go.uber.org/zap"
)

type EmailService interface {
	SendEmail(to, subject, body string) error
}

type emailService struct {
	logger *zap.Logger
}

func NewEmailService(logger *zap.Logger) EmailService {
	return &emailService{
		logger: logger.Named("EmailService"),
	}
}

func (s *emailService) SendEmail(to, subject, body string) error {
	sugar := s.logger.Sugar()

	sugar.Infow("Simulating email send",
		"to", to,
		"subject", subject,
	)

	return nil
}
