package object_values

import (
	"github.com/cassiusbessa/backend-test/internal/domain/shared"
	"golang.org/x/crypto/bcrypt"
)

type Password struct {
	hash string
}

func NewPassword(plainText string) (Password, error) {
	if len(plainText) < 6 {
		return Password{}, shared.ErrValidation
	}

	hashed, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		return Password{}, err
	}

	return Password{hash: string(hashed)}, nil
}

func (p Password) Hash() string {
	return p.hash
}

func (p Password) Compare(plainText string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(p.hash), []byte(plainText))
	return err == nil
}
