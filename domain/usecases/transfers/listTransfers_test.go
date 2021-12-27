package transfers

import (
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	accUse "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/google/uuid"
)

func TestTransfers(t *testing.T) {
	person := account.Account{
		Id:      uuid.New().String(),
		Name:    "Vanny",
		Cpf:     "55566689545",
		Secret:  "dafd33255",
		Balance: 2500,
	}

	person2 := account.Account{
		Id:      uuid.New().String(),
		Name:    "Viviane",
		Cpf:     "11452369875",
		Secret:  "vivi",
		Balance: 2500,
	}

	t.Run("Should list a transfer successfull", func(t *testing.T) {
		transMock := transfers.TransferMock{
			OnListAll: func(id string) ([]transfers.Transfer, error) {
				return []transfers.Transfer{
					{AccountOriginId: person.Id,
						AccountDestinationId: person2.Id,
						Amount:               100},
				}, nil
			},
		}

		accMock := account.AccountMock{
			OnCreate: func(account account.Account) (account.Account, error) {
				return person, nil
			},
		}

		accUsecase := accUse.CreateNewAccountUsecase(accMock)

		transUsecase := CreateNewTransferUsecase(transMock, accUsecase)

		list, err := transUsecase.repo.ListTransfer(person.Id)

		if err != nil {
			t.Errorf("Expected error to be nil, got %s", err)
		}

		if len(list) != 1 {
			t.Errorf("Expected 1 item in list transfer, got %v", len(list))
		}
	})
}
