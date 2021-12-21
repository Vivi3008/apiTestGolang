package usecases

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/store"
)

func TestTransfers(t *testing.T) {
	transStore := store.NewTransferStore()
	trans := SaveNewTransfer(transStore)

	t.Run("Should create a transfer, save and list it successfull", func(t *testing.T) {
		accountStore := store.NewAccountStore()
		accounts := CreateAccountStore(accountStore)

		person := account.Account{
			Name:    "Vanny",
			Cpf:     "55566689545",
			Secret:  "dafd33255",
			Balance: 2500,
		}

		person2 := account.Account{
			Name:    "Viviane",
			Cpf:     "11452369875",
			Secret:  "vivi",
			Balance: 2500,
		}

		acc, err := account.NewAccount(person)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		account2, err := account.NewAccount(person2)

		if err != nil {
			t.Errorf("Expected nil, got %s", err)
		}

		transfer := transfers.Transfer{
			AccountOriginId:      acc.Id,
			AccountDestinationId: account2.Id,
			Amount:               5,
		}

		transOk, err := accounts.CreateTransfer(transfer)

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if transOk.CreatedAt.IsZero() {
			t.Errorf("Expected createdAt at not to be zero")
		}

		_, err = trans.SaveTransfer(transOk)

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		listTransfers, err := trans.ListTransfer(transOk.AccountOriginId)

		if err != nil {
			t.Errorf("expected nil, got %s", err.Error())
		}

		if listTransfers[0].AccountOriginId != transOk.AccountOriginId {
			t.Errorf("expected %s; got %s", listTransfers[0].AccountOriginId, transOk.AccountOriginId)
		}
	})

	t.Run("Should not list transfers if id origin doesnt exists", func(t *testing.T) {
		_, err := trans.ListTransfer("f63cb25b-786c-4ff2-9a67-22a065d307d3")

		if err == nil {
			t.Errorf("Expected err origin id doesnt exists, got %s", err)
		}
	})
}
