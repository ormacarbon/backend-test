package services

import (
	gomail "gopkg.in/mail.v2"
)

// Instantiate a new EmailService
func NewEmailService(emailConfig EmailConfig) *EmailService {
	return &EmailService{
		EmailConfig: emailConfig,
	}
}

// Send a welcome email to the user upon registration
func (s *EmailService) SendWelcomeEmail(email string) error {
	// Creating a new message
	message := gomail.NewMessage()

	// Setting email headers
	message.SetHeader("From", s.EmailConfig.SMTPEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Welcome to GSS Eco News!")

	// Setting email body
	message.SetBody("text/plain", "Welcome to GSS Eco News! We are excited to have you on board!")

	// Setting up SMTP configuration
	dialer := gomail.NewDialer(
		s.EmailConfig.SMTPHost,
		s.EmailConfig.SMTPPort,
		s.EmailConfig.SMTPEmail,
		s.EmailConfig.SMTPPassword,
	)

	// Sending the email
	return dialer.DialAndSend(message)
}

// Send email to user when someones is registered using their referral link
func (s *EmailService) SendReferralLinkAccess(email string) error {
	// Creating a new message
	message := gomail.NewMessage()

	// Setting email headers
	message.SetHeader("From", s.EmailConfig.SMTPEmail)
	message.SetHeader("To", email)
	message.SetHeader("Subject", "Someone has registered using your referral link!")

	// Setting email body
	message.SetBody("text/plain", "Someone has registered using your referral link! You have earned 1 point in the competition!")

	// Setting up SMTP configuration
	dialer := gomail.NewDialer(
		s.EmailConfig.SMTPHost,
		s.EmailConfig.SMTPPort,
		s.EmailConfig.SMTPEmail,
		s.EmailConfig.SMTPPassword,
	)

	// Sending the email
	return dialer.DialAndSend(message)
}