package account

type AccountMock struct {
	OnStoreAccount func(account Account) error
	OnListAll      func() ([]Account, error)
	OnListById     func(accountId string) (Account, error)
}

func (m AccountMock) StoreAccount(account Account) error {
	return m.OnStoreAccount(account)
}

func (m AccountMock) ListAllAccounts() ([]Account, error) {
	return m.OnListAll()
}

func (m AccountMock) ListAccountById(accountId string) (Account, error) {
	return m.OnListById(accountId)
}
