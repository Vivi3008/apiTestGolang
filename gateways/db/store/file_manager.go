package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type Entity string

const accountType Entity = "account"

type Entities struct {
	Account  []account.Account
	Transfer map[string]transfers.Transfer
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
	var transferData map[string]transfers.Transfer

	dataJson, err := ioutil.ReadFile(source)

	if err != nil {
		return Entities{}, err
	}

	if typeEntitie == accountType {
		json.Unmarshal(dataJson, &accountData)
		return Entities{Account: accountData}, nil
	} else {
		json.Unmarshal(dataJson, &transferData)
		return Entities{Transfer: transferData}, nil
	}
}
