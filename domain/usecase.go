package domain

import (
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type Usecase interface {
	CreateAccount(account.Account) (account.Account, error)
	ListAllAccounts() ([]account.Account, error)
	ListAccountById(string) (account.Account, error)
	NewLogin(account.Login) (string, error)
	ListTransfer(string) (transfers.Transfer, error)
	SaveTransfer(transfers.Transfer) (transfers.Transfer, error)
	CreateTransfer(transfers.Transfer) (transfers.Transfer, error)
	CreateBill(bills.Bill) (bills.Bill, error)
	SaveBill(bills.Bill) error
	ListAllBills(string) ([]bills.Bill, error)
}
