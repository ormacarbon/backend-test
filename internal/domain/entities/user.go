package entities

import (
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
)

type User struct {
	name     string
	email    object_values.Email
	phone    object_values.PhoneNumber
	password object_values.Password
}

func NewUser(name string, email object_values.Email, hashed_pass object_values.Password, phone object_values.PhoneNumber) (User, error) {
	if name == "" {
		return User{}, shared.ErrValidation
	}
	return User{name: name, password: hashed_pass, email: email, phone: phone}, nil
}

func (u User) Name() string {
	return u.name
}

func (u User) Email() object_values.Email {
	return u.email
}

func (u User) Password() object_values.Password {
	return u.password
}

func (u User) Phone() object_values.PhoneNumber {
	return u.phone
}
