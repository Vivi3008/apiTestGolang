package transfers

import (
	"context"
	"errors"
	"sort"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/db/store"
)

var (
	ErrIdNotExists  = errors.New("id does not exist")
	ErrCpfNotExists = errors.New("cpf does not exist")
)

func (t TransferStore) ListTransfer(ctx context.Context, id string) ([]transfers.Transfer, error) {
	data, err := store.ReadFile(t.Src, "transfer")
	if err != nil {
		return []transfers.Transfer{}, err
	}

	for _, tr := range data.Transfer {
		if tr.AccountOriginId == id {
			t.tranStore = append(t.tranStore, tr)
		}
	}

	t.OrderListTransferByDateDesc()

	return t.tranStore, nil
}

func (t TransferStore) OrderListTransferByDateDesc() {
	sort.Slice(t.tranStore, func(i, j int) bool {
		return t.tranStore[i].CreatedAt.After(t.tranStore[j].CreatedAt)
	})
}
