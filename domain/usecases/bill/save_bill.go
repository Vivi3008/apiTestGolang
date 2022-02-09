package bill

import (
	"context"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

func (b BillUsecase) SaveBill(ctx context.Context, bill bills.Bill) error {
	err := b.blRepo.StoreBill(ctx, bill)
	if err != nil {
		return err
	}
	return nil
}
