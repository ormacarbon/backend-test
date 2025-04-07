package internal

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAuthor(t *testing.T) {
	t.Run("should return an error if the name is invalid", func(t *testing.T) {
		author, err := NewAuthor("", "test@example.com", "5531999999999")
		assert.Nil(t, author)
		assert.Equal(t, "invalid name", err.Error())
	})
	t.Run("should return an error if the email is invalid", func(t *testing.T) {
		author, err := NewAuthor("John Doe", "invalid-email", "5511987654321")
		assert.Nil(t, author)
		assert.Equal(t, "invalid email", err.Error())
	})
	t.Run("should return an error if the phone is invalid", func(t *testing.T) {
		author, err := NewAuthor("John Doe", "johndoe@email.com", "123456789")
		assert.Nil(t, author)
		assert.Equal(t, "invalid phone", err.Error())
	})
}
