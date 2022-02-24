package activity

import (
	"io/ioutil"
	"os"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
	"github.com/google/uuid"
)

var sourceTest = "../bills/bills_test.json"

var accountId = uuid.NewString()

var BillsTest = []bills.Bill{
	{
		Id:            uuid.New().String(),
		AccountId:     accountId,
		Description:   "Academia",
		Value:         5990,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
	},
	{
		Id:            uuid.New().String(),
		AccountId:     accountId,
		Description:   "Internet",
		Value:         15000,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
	},
	{
		Id:            uuid.New().String(),
		AccountId:     accountId,
		Description:   "IPTU",
		Value:         130000,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
	},
}

func CreateBillsActivities() error {
	err := store.StoreFile(BillsTest, sourceTest)
	if err != nil {
		return err
	}
	return nil
}

func DeleteBillsActivities() error {
	err := ioutil.WriteFile(sourceTest, []byte{}, os.ModeAppend)
	if err != nil {
		return err
	}
	return nil
}
