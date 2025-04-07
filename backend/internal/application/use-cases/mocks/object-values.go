package mocks

import (
	"testing"

	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/stretchr/testify/assert"
)

func MustEmail(t *testing.T, value string) object_values.Email {
	email, err := object_values.NewEmail(value)
	assert.NoError(t, err)
	return email
}

func MustPassword(t *testing.T, value string) object_values.Password {
	pass, err := object_values.NewPassword(value)
	assert.NoError(t, err)
	return pass
}

func MustPhone(t *testing.T, value string) object_values.PhoneNumber {
	phone, err := object_values.NewPhoneNumber(value)
	assert.NoError(t, err)
	return phone
}
