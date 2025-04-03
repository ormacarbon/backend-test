package object_values

import (
	"regexp"

	"github.com/cassiusbessa/backend-test/internal/domain/shared"
)

type PhoneNumber struct {
	value string
}

var phoneRegex = regexp.MustCompile(`^\+[1-9]\d{7,14}$`)

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	if !phoneRegex.MatchString(phone) {
		return PhoneNumber{}, shared.ErrValidation
	}
	return PhoneNumber{value: phone}, nil
}

func (p PhoneNumber) Value() string {
	return p.value
}
