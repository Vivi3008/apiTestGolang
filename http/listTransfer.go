package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain"
)

type TransferResponse struct {
	Id                   string `json:"id"`
	AccountOriginId      string `json:"account_origin_id"`
	AccountDestinationId string `json:"account_destination_id"`
	Amount               int64  `json:"amount"`
	CreatedAt            string `json:"createdAt"`
}

func (s Server) ListTransfer(w http.ResponseWriter, r *http.Request) {
	accountId, _ := VerifyAuth(w, r)

	list, err := s.tr.ListTransfer(domain.AccountId(accountId))

	if err != nil {
		log.Printf("Failed to do list transfers: %s\n", err.Error())
		return
	}

	response := make([]TransferResponse, len(list))

	for i, transfer := range list {
		response[i].Id = transfer.Id
		response[i].AccountOriginId = transfer.AccountOriginId
		response[i].AccountDestinationId = transfer.AccountDestinationId
		response[i].Amount = transfer.Amount
		response[i].CreatedAt = transfer.CreatedAt.Format(DateLayout)
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("Sent all transfers from Id %s", accountId)
	log.Printf("Sent all transfers. Total: %d", len(response))
}
