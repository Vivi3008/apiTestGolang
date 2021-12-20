package account

type AccountRepository interface {
	StoreAccount(account Account) error
	ListAllAccounts() ([]Account, error)
	ListAccountById(accountId string) (Account, error)
}
