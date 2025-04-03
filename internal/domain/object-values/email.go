package object_values

import (
	"errors"
	"regexp"
)

type Email struct {
	value string
}

var (
	emailRegex      = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	errInvalidEmail = errors.New("invalid email format")
)

func NewEmail(email string) (Email, error) {
	if !emailRegex.MatchString(email) {
		return Email{}, errInvalidEmail
	}
	return Email{value: email}, nil
}

func (e Email) Value() string {
	return e.value
}
