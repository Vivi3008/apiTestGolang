package bill

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
)

//cria o pagamento e atualiza a conta
func (b BillUsecase) CreateBill(ctx context.Context, bill bills.Bill) (bills.Bill, error) {
	pay, err := bills.NewBill(bill)

	if err != nil {
		return bills.Bill{}, err
	}

	_, err = b.acRepo.UpdateAccountBalance(ctx, bill.AccountId, bill.Value, account.Debit)

	if err != nil {
		return bills.Bill{}, err
	}

	return pay, nil
}
