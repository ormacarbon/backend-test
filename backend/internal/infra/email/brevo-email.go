package email

import (
	"context"
	"os"

	output_ports "github.com/cassiusbessa/backend-test/internal/application/ports/output"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/getbrevo/brevo-go/lib"
)

type BrevoEmailService struct {
	client *lib.APIClient
	sender lib.SendSmtpEmailSender
}

func NewBrevoEmailService(senderEmail, senderName string) output_ports.EmailService {
	cfg := lib.NewConfiguration()
	cfg.AddDefaultHeader("api-key", os.Getenv("BREVO_API_KEY"))

	client := lib.NewAPIClient(cfg)

	return &BrevoEmailService{
		client: client,
		sender: lib.SendSmtpEmailSender{Email: senderEmail, Name: senderName},
	}
}

func (s *BrevoEmailService) SendEmail(to string, subject string, body string) error {
	ctx := context.Background()

	email := lib.SendSmtpEmail{
		Sender:      &s.sender,
		To:          []lib.SendSmtpEmailTo{{Email: to}},
		Subject:     subject,
		TextContent: body,
	}

	_, _, err := s.client.TransactionalEmailsApi.SendTransacEmail(ctx, email)
	if err != nil {
		return shared.ErrInternal
	}

	return nil
}
