package internal

import (
	"fmt"
	"log"
	"net/smtp"
)

type SendEmail interface {
	Send(input InputSendEmail) error
}

type SMTP struct {
	host     string
	port     int
	username string
	password string
}

type InputSendEmail struct {
	To      string
	Subject string
	Body    string
}

func NewSMTP(host string, port int, user string, pass string) *SMTP {
	return &SMTP{
		host:     host,
		port:     port,
		username: user,
		password: pass,
	}
}

func (s *SMTP) Send(input InputSendEmail) error {
	auth := smtp.PlainAuth("", s.username, s.password, s.host)
	to := []string{input.To}
	msg := []byte(input.Body)
	addr := fmt.Sprintf("%s:%d", s.host, s.port)
	log.Printf("host: %s, username: %s, password: %s", addr, s.username, s.password)
	log.Printf("to: %s, subject: %s, body: %s", input.To, input.Subject, input.Body)
	err := smtp.SendMail(addr, auth, "equipe@carbon_offsets.com", to, msg)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}
