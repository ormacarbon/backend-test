package services


type IEmailService interface {
	SendWelcomeEmail(email string) error
	SendReferralLinkAccess(email string) error
	SendLeaderboardEmail(email string) error
}

type EmailConfig struct {
	SMTPEmail string
	SMTPHost    string
	SMTPPort    int
	SMTPPassword   string
}

type EmailService struct {
	EmailConfig EmailConfig
}
