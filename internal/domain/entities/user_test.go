package entities_test

import (
	"testing"

	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
)

func TestNewUser(t *testing.T) {

	tests := []struct {
		name        string
		email       string
		phone       string
		expectError bool
	}{
		{"John Doe", "user@example.com", "+14155552671", false},
		{"", "user@example.com", "+14155552671", true},
	}

	for _, tt := range tests {
		email, _ := object_values.NewEmail(tt.email)
		phone, _ := object_values.NewPhoneNumber(tt.phone)

		_, err := entities.NewUser(tt.name, email, phone)
		if (err != nil) != tt.expectError {
			t.Errorf("NewUser(%s, %s, %s) expected error: %v, got: %v", tt.name, tt.email, tt.phone, tt.expectError, err)
		}
	}
}
