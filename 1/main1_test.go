package main

import (
	"crypto/sha256"
	"encoding/hex"
	"testing"
)

func TestProcessVariables(t *testing.T) {
	tests := []struct {
		name         string
		input        []interface{}
		expected     string
		expectedHash string
	}{
		{
			name:     "Тест с integer, string и float64",
			input:    []interface{}{42, "test", 3.14},
			expected: "42 test 3.14 ",
		},
		{
			name:     "Тест с  bool и complex",
			input:    []interface{}{true, false, complex(1, 2)},
			expected: "true false (1+2i) ",
		},
	}

	salt := "go-2024"

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			combinedString, hash := processVariables(tt.input)

			if combinedString != tt.expected {
				t.Errorf("Expected combined string %q, got %q", tt.expected, combinedString)
			}

			// Генерация ожидаемого хэша
			runes := []rune(tt.expected)
			saltedRunes := append(runes[:len(runes)/2], []rune(salt)...)
			saltedRunes = append(saltedRunes, runes[len(runes)/2:]...)
			expectedHash := sha256.Sum256([]byte(string(saltedRunes)))
			expectedHashStr := hex.EncodeToString(expectedHash[:])

			if hash != expectedHashStr {
				t.Errorf("Expected hash %q, got %q", expectedHashStr, hash)
			}
		})
	}
}
