package email

import (
	"bytes"
	"fmt"
	"html/template"
	"net/smtp"

	"github.com/icl00ud/backend-test/internal/config"
	"go.uber.org/zap"
)

type Service interface {
	Email(to, subject, templatePath string, data interface{})
}

type emailService struct {
	smtpHost     string
	smtpPort     string
	smtpUsername string
	smtpPassword string
	logger       *zap.SugaredLogger
}

func NewEmailService(cfg *config.Config, logger *zap.SugaredLogger) Service {
	return &emailService{
		smtpHost:     cfg.SMTPHost,
		smtpPort:     cfg.SMTPPort,
		smtpUsername: cfg.SMTPUsername,
		smtpPassword: cfg.SMTPPassword,
		logger:       logger.Named("EmailService"),
	}
}

func (s *emailService) Email(to, subject, templatePath string, data interface{}) {
	tmpl, err := template.ParseFiles(templatePath)
	if err != nil {
		s.logger.Errorw("Failed to parse HTML template", "error", err)
		return
	}

	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		s.logger.Errorw("Failed to execute HTML template", "error", err)
		return
	}

	htmlBody := buf.String()
	from := s.smtpUsername

	msg := []byte("From: " + from + "\r\n" +
		"To: " + to + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"UTF-8\"\r\n\r\n" +
		htmlBody + "\r\n",
	)
	addr := fmt.Sprintf("%s:%s", s.smtpHost, s.smtpPort)
	auth := smtp.PlainAuth("", s.smtpUsername, s.smtpPassword, s.smtpHost)
	s.logger.Infow("Sending email", "to", to, "subject", subject)
	if err := smtp.SendMail(addr, auth, from, []string{to}, msg); err != nil {
		s.logger.Errorw("Failed to send email", "error", err)
		return
	}
}
