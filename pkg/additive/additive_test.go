package additive

import (
	"testing"
)

func TestAdditiveCipher(t *testing.T) {
	t.Run("Encrypt 1", func(t *testing.T) {
		plaintext := "hello"
		key := 3
		expected := "khoor"
		actual := Encrypt(plaintext, key)
		if actual != expected {
			t.Errorf("Expected %s but got %s", expected, actual)
		}

	})

	t.Run("Encrypt 2", func(t *testing.T) {
		plaintext := "world"
		key := 5
		expected := "btwqi"
		actual := Encrypt(plaintext, key)
		if actual != expected {
			t.Errorf("Expected %s but got %s", expected, actual)
		}
	})
}
