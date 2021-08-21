package domain

import (
	"testing"
)

func TestStoreAccount(t *testing.T) {

	t.Run("Should create an account successfully", func(t *testing.T) {
		person := Account{
			name: "Vanny",
			cpf:     13323332555,
			secret:  "dafd33255",
			balance: 2.500,
		}

		result, err := NewAccount(person)
		expected := "Vanny"

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if result.name != expected {
			t.Errorf("Resultado %v, esperado %v", result.name, expected)
		}

		if result.createdAt.IsZero(){
			t.Errorf("Expected createdAt at not to be zero")
		}

		if result.id == ""{
			t.Errorf("Expected ID not to be empty")
		}
	})
}
