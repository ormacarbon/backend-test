package entities_test

import (
	"testing"

	"github.com/cassiusbessa/backend-test/internal/domain/entities"
	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
	"github.com/google/uuid"
)

func TestNewUser(t *testing.T) {

	tests := []struct {
		name        string
		email       string
		phone       string
		invitedBy   *uuid.UUID
		expectError bool
	}{
		{"John Doe", "user@example.com", "+14155552671", nil, false},
		{"", "user@example.com", "+14155552671", nil, true},
	}

	for _, tt := range tests {
		email, _ := object_values.NewEmail(tt.email)
		phone, _ := object_values.NewPhoneNumber(tt.phone)
		password, _ := object_values.NewPassword("defaultPassword123")

		u, err := entities.NewUser(tt.name, email, password, phone, tt.invitedBy)
		if (err != nil) != tt.expectError {
			t.Errorf("NewUser(%s, %s, %s) expected error: %v, got: %v", tt.name, tt.email, tt.phone, tt.expectError, err)
		}

		if err == nil && u.Points() != 1 {
			t.Errorf("expected initial points to be 1, got %d", u.Points())
		}
	}
}
