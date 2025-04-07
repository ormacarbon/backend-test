package object_values

import (
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
)

type PhoneNumber struct {
	value string
}

func NewPhoneNumber(phone string) (PhoneNumber, error) {
	if len(phone) < 8 || len(phone) > 20 {
		return PhoneNumber{}, shared.ErrValidation
	}
	return PhoneNumber{value: phone}, nil
}

func (p PhoneNumber) Value() string {
	return p.value
}
