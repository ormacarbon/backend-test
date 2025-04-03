package object_values_test

import (
	"testing"

	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
)

func TestNewEmail(t *testing.T) {
	tests := []struct {
		email       string
		expectError bool
	}{
		{"test@example.com", false},
		{"invalidemail", true},
		{"@nouser.com", true},
		{"nodomain@", true},
		{"", true},
	}

	for _, tt := range tests {
		_, err := object_values.NewEmail(tt.email)
		if (err != nil) != tt.expectError {
			t.Errorf("NewEmail(%s) expected error: %v, got: %v", tt.email, tt.expectError, err)
		}
	}
}
