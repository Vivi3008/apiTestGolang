package store

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type Entity string

const (
	accountType  Entity = "account"
	billType     Entity = "bill"
	transferType Entity = "transfer"
)

type Entities struct {
	Account  []account.Account
	Bill     []bills.Bill
	Transfer []transfers.Transfer
}

func StoreFile(writeData interface{}, source string) error {
	data, err := json.Marshal(writeData)

	if err != nil {
		return err
	}

	err = os.Chmod(source, 0777)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(source, data, os.ModeAppend)

	if err != nil {
		return err
	}

	return nil
}

func ReadFile(source string, typeEntitie Entity) (Entities, error) {
	var accountData []account.Account
	var billData []bills.Bill
	var transferData []transfers.Transfer

	err := os.Chmod(source, 0777)
	if err != nil {
		return Entities{}, err
	}
	dataJson, err := ioutil.ReadFile(source)

	if err != nil {
		return Entities{}, err
	}

	if !json.Valid(dataJson) {
		return Entities{}, nil
	}

	switch typeEntitie {
	case accountType:
		err := json.Unmarshal(dataJson, &accountData)
		return sendError(err, Entities{Account: accountData})
	case billType:
		err := json.Unmarshal(dataJson, &billData)
		return sendError(err, Entities{Bill: billData})
	case transferType:
		err := json.Unmarshal(dataJson, &transferData)
		return sendError(err, Entities{Transfer: transferData})
	default:
		return Entities{}, nil
	}
}

func sendError(err error, entity Entities) (Entities, error) {
	if err != nil {
		return Entities{}, err
	}
	return entity, nil
}
