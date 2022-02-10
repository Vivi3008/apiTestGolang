package bill

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b BillUsecase) ListAllBills(ctx context.Context, accountId string) ([]bills.Bill, error) {
	_, err := b.acRepo.ListAccountById(ctx, accountId)

	if err != nil {
		return []bills.Bill{}, err
	}

	list, err := b.blRepo.ListBills(ctx, accountId)

	if err != nil {
		return []bills.Bill{}, err
	}

	return list, nil
}
