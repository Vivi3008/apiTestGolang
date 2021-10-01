package domain

type Usecase interface {
	CreateAccount(Account) (Account, error)
	ListAllAccounts() ([]Account, error)
	ListAccountById(AccountId) (Account, error)
}
