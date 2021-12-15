package domain

type Usecase interface {
	CreateAccount(Account) (Account, error)
	ListAllAccounts() ([]Account, error)
	ListAccountById(string) (Account, error)
	NewLogin(Login) (string, error)
	ListTransfer(string) (Transfer, error)
	SaveTransfer(Transfer) (Transfer, error)
	CreateTransfer(Transfer) (Transfer, error)
	CreateBill(Bill) (Bill, error)
	SaveBill(Bill) (Bill, error)
	ListAllBills(string) ([]Bill, error)
}
