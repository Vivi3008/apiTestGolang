package store

import (
	"encoding/json"
	"errors"
	"fmt"
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

var ErrSaveInFile = errors.New("error to save in file")

func StoreFile(writeData interface{}, source string) error {
	data, err := json.Marshal(writeData)

	if err != nil {
		return fmt.Errorf(ErrSaveInFile.Error(), err)
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

	switch typeEntitie {
	case accountType:
		json.Unmarshal(dataJson, &accountData)
		return Entities{Account: accountData}, nil
	case billType:
		json.Unmarshal(dataJson, &billData)
		return Entities{Bill: billData}, nil
	case transferType:
		json.Unmarshal(dataJson, &transferData)
		return Entities{Transfer: transferData}, nil
	default:
		return Entities{}, nil
	}
}
