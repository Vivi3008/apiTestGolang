package domain

import (
	"testing"
)

func TestStoreAccount(t *testing.T) {

	t.Run("Should create an account successfully", func(t *testing.T) {
		person := Account{
			Name:    "Vanny",
			Cpf:     13323332555,
			Secret:  "dafd33255",
			Balance: 2.500,
		}

		result, err := NewAccount(person)
		expected := "Vanny"

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if result.Name != expected {
			t.Errorf("Resultado %v, esperado %v", result.Name, expected)
		}

		if result.createdAt.IsZero(){
			t.Errorf("Expected createdAt at not to be zero")
		}

		if result.Id == ""{
			t.Errorf("Expected ID not to be empty")
		}
	})
}
