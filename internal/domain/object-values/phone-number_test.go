package object_values_test

import (
	"testing"

	object_values "github.com/cassiusbessa/backend-test/internal/domain/object-values"
)

func TestNewPhoneNumber(t *testing.T) {
	tests := []struct {
		phone       string
		expectError bool
	}{
		{"+14155552671", false},   // Válido: EUA
		{"+5511987654321", false}, // Válido: Brasil
		{"+442071838750", false},  // Válido: Reino Unido
		{"+123", true},            // Inválido: Muito curto
	}

	for _, tt := range tests {
		_, err := object_values.NewPhoneNumber(tt.phone)
		if (err != nil) != tt.expectError {
			t.Errorf("NewPhoneNumber(%s) expected error: %v, got: %v", tt.phone, tt.expectError, err)
		}
	}
}
