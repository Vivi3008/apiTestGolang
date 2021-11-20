package domain

type Usecase interface {
	CreateAccount(Account) (Account, error)
	ListAllAccounts() ([]Account, error)
	ListAccountById(AccountId) (Account, error)
	NewLogin(Login) (AccountId, error)
	ListTransfer(AccountId) (Transfer, error)
	SaveTransfer(Transfer) (Transfer, error)
	CreateTransfer(Transfer) (Transfer, error)
	CreateBill(Bill) (Bill, error)
	SaveBill(Bill) (Bill, error)
	ListAllBills(AccountId) ([]Bill, error)
}
