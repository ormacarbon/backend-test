package internal

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

const (
	minNameLength = 3
	maxNameLength = 60
)

var (
	InvalidName  = errors.New("invalid name")
	InvalidEmail = errors.New("invalid email")
	InvalidPhone = errors.New("invalid phone")
)

type Author struct {
	ID           int
	Name         string
	Email        string
	Phone        string
	Points       uint8
	ReferralCode string
	CreatedAt    time.Time
}

func NewAuthor(name, email, phone string) (*Author, error) {
	author := Author{
		Name:         name,
		Email:        email,
		Phone:        phone,
		Points:       1,
		ReferralCode: strings.ToLower(fmt.Sprintf("@%s", strings.ReplaceAll(name, " ", ""))),
	}
	if author.validateName() {
		return nil, InvalidName
	}
	if !author.validateEmail() {
		return nil, InvalidEmail
	}
	if !author.validatePhone() {
		return nil, InvalidPhone
	}
	return &author, nil
}

func (u *Author) validateName() bool {
	if len(u.Name) < minNameLength || len(u.Name) > maxNameLength {
		return true
	}
	return false
}

func (u *Author) validateEmail() bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(u.Email)
}

func (u *Author) validatePhone() bool {
	phoneRegex := regexp.MustCompile(`^(?:(?:\+|00)?(55)\s?)?(?:\(?([1-9][0-9])\)?\s?)?(?:((?:9\d|[2-9])\d{3})\-?(\d{4}))$`)
	return phoneRegex.MatchString(u.Phone)
}
