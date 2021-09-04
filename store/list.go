package store

import "github.com/Vivi3008/apiTestGolang/domain"

func (a AccountStore) ListAll()([]domain.Account, error){
	var list []domain.Account
	for _, account := range a.accStore {
		list = append(list, account)
	}
	return list, nil
}

// fazer o teste de listar um id
func (a AccountStore) ListOne(accId string) (domain.Account, error){
	var listOne domain.Account
	listOne = a.accStore[accId]

	return listOne, nil
}
