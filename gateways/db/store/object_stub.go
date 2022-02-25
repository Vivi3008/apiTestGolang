package store

import (
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/google/uuid"
)

var AccountsTest = []account.Account{
	{
		Id:        uuid.NewString(),
		Name:      "Teste 1",
		Cpf:       "77845100032",
		Secret:    "dafd33255",
		Balance:   250000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste 2",
		Cpf:       "55985633301",
		Secret:    "4f5ds4af54",
		Balance:   260000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste 3",
		Cpf:       "85665232145",
		Secret:    "fadsfdsaf",
		Balance:   360000,
		CreatedAt: time.Now(),
	},
}

var accountId = "0e4a7fdd-59bd-4f99-817b-225c732b11f9"

var BillsTest = []bills.Bill{
	{
		Id:            uuid.New().String(),
		AccountId:     accountId,
		Description:   "Academia",
		Value:         5990,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
		CreatedAt: time.Date(2022, time.February,
			8, 21, 34, 01, 0, time.UTC),
	},
	{
		Id:            uuid.New().String(),
		AccountId:     AccountsTest[0].Id,
		Description:   "Internet",
		Value:         15000,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
		CreatedAt: time.Date(2022, time.February,
			9, 21, 34, 01, 0, time.UTC),
	},
	{
		Id:            uuid.New().String(),
		AccountId:     AccountsTest[0].Id,
		Description:   "IPTU",
		Value:         130000,
		DueDate:       time.Now().AddDate(0, 0, 2),
		ScheduledDate: time.Now(),
		CreatedAt: time.Date(2022, time.February,
			10, 21, 34, 01, 0, time.UTC),
	},
}

var TransfersTest = []transfers.Transfer{
	{
		Id:                   uuid.NewString(),
		AccountOriginId:      AccountsTest[1].Id,
		AccountDestinationId: AccountsTest[0].Id,
		Amount:               100000,
		CreatedAt: time.Date(2022, time.February,
			8, 21, 34, 01, 0, time.UTC),
	},
	{
		Id:                   uuid.NewString(),
		AccountOriginId:      AccountsTest[0].Id,
		AccountDestinationId: AccountsTest[2].Id,
		Amount:               200000,
		CreatedAt: time.Date(2022, time.February,
			9, 21, 34, 01, 0, time.UTC),
	},
	{
		Id:                   uuid.NewString(),
		AccountOriginId:      AccountsTest[0].Id,
		AccountDestinationId: AccountsTest[1].Id,
		Amount:               300000,
		CreatedAt: time.Date(2022, time.February,
			10, 21, 34, 01, 0, time.UTC),
	},
}
