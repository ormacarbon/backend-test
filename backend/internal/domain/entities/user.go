package entities

import (
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"github.com/google/uuid"
)

type User struct {
	id         uuid.UUID
	name       string
	email      object_values.Email
	phone      object_values.PhoneNumber
	password   object_values.Password
	inviteCode string
	invitedBy  *uuid.UUID
	points     int
}

func NewUser(
	name string,
	email object_values.Email,
	hashedPass object_values.Password,
	phone object_values.PhoneNumber,
	invitedBy *uuid.UUID,
) (User, error) {
	if name == "" {
		return User{}, shared.ErrValidation
	}

	return User{
		id:         uuid.New(),
		name:       name,
		email:      email,
		phone:      phone,
		password:   hashedPass,
		inviteCode: uuid.NewString(),
		invitedBy:  invitedBy,
		points:     1,
	}, nil
}

func LoadUser(
	id uuid.UUID,
	name string,
	email object_values.Email,
	hashedPass object_values.Password,
	phone object_values.PhoneNumber,
	inviteCode string,
	invitedBy *uuid.UUID,
	points int,
) User {
	return User{
		id:         id,
		name:       name,
		email:      email,
		phone:      phone,
		password:   hashedPass,
		inviteCode: inviteCode,
		invitedBy:  invitedBy,
		points:     points,
	}
}

func (u User) ID() uuid.UUID {
	return u.id
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

func (u User) InviteCode() string {
	return u.inviteCode
}

func (u User) InvitedBy() *uuid.UUID {
	return u.invitedBy
}

func (u *User) AddPoint() {
	u.points++
}

func (u User) Points() int {
	return u.points
}
