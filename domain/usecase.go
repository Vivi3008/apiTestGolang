package domain

import "github.com/Vivi3008/apiTestGolang/domain/entities/account"

type Usecase interface {
	CreateAccount(account.Account) (account.Account, error)
	ListAllAccounts() ([]account.Account, error)
	ListAccountById(string) (account.Account, error)
	NewLogin(account.Login) (string, error)
	ListTransfer(string) (account.Transfer, error)
	SaveTransfer(account.Transfer) (account.Transfer, error)
	CreateTransfer(account.Transfer) (account.Transfer, error)
	CreateBill(Bill) (Bill, error)
	SaveBill(Bill) (Bill, error)
	ListAllBills(string) ([]Bill, error)
}
