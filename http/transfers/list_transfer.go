package transfers

import (
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/http/response"
)

type TransferResponse struct {
	Id                   string `json:"id"`
	AccountOriginId      string `json:"account_origin_id"`
	AccountDestinationId string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"createdAt"`
}

func (h Handler) ListTransfer(w http.ResponseWriter, r *http.Request) {
	accountId, ok := middlewares.GetAccountId(r.Context())

	if !ok || accountId == "" {
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	account, err := h.accUse.ListAccountById(accountId)

	if err != nil {
		log.Printf("Failed to list transfer: %s\n", err.Error())
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	list, err := h.transfUse.ListTransfer(account.Id)

	if err != nil {
		log.Printf("Failed to list transfer: %s\n", err.Error())
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	transfers := make([]TransferResponse, len(list))

	for i, transfer := range list {
		transfers[i].Id = transfer.Id
		transfers[i].AccountOriginId = transfer.AccountOriginId
		transfers[i].AccountDestinationId = transfer.AccountDestinationId
		transfers[i].Amount = transfer.Amount
		transfers[i].CreatedAt = transfer.CreatedAt.Format(response.DateLayout)
	}

	response.SendRequest(w, transfers, http.StatusOK)
	log.Printf("Sent all transfers from Id %s", accountId)
	log.Printf("Sent all transfers. Total: %d", len(transfers))
}
