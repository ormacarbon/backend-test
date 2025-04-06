package output_ports

type EmailService interface {
	SendEmail(to string, subject string, body string) error
}
