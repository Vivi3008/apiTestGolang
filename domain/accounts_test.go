package domain

import (
	"testing"
)

func TestStoreAccount(t *testing.T) {

	t.Run("Should store an account successfuly", func(t *testing.T) {
		person := Account{
			name:    "Vanny",
			cpf:     13323332555,
			secret:  "dafd33255",
			balance: 2.500,
		}
		result, err := NewAccount(person)
		expected := Account{
			name:    "Viviane",
			cpf:     13323332555,
			secret:  "dafd33255",
			balance: 2.500,
		}

		if result != expected {
			t.Errorf("Resultado %T, esperado %T, erro %s", result, expected, err)
		}
	})
}
