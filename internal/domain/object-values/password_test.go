package object_values_test

import (
	"testing"

	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
)

func TestNewPassword(t *testing.T) {
	tests := []struct {
		password    string
		expectError bool
	}{
		{"password123", false},
		{"short", true},
		{"", true},
	}

	for _, tt := range tests {
		_, err := object_values.NewPassword(tt.password)
		if (err != nil) != tt.expectError {
			t.Errorf("NewPassword(%q) expected error: %v, got: %v", tt.password, tt.expectError, err)
		}
	}
}

func TestPasswordCompare(t *testing.T) {
	pass, _ := object_values.NewPassword("password123")

	if !pass.Compare("password123") {
		t.Errorf("Expected password to match")
	}

	if pass.Compare("wrongpassword") {
		t.Errorf("Expected password to NOT match")
	}
}
