package object_values

import (
	"regexp"

	"github.com/cassiusbessa/backend-test/internal/domain/shared"
)

type Email struct {
	value string
}

var (
	emailRegex = regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
)

func NewEmail(email string) (Email, error) {
	if !emailRegex.MatchString(email) {
		return Email{}, shared.ErrValidation
	}
	return Email{value: email}, nil
}

func (e Email) Value() string {
	return e.value
}
