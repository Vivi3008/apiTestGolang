package transfers

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/http/response"
)

type TransferRequest struct {
	AccountDestinationId string `json:"account_Destination_Id"`
	Amount               int    `json:"amount"`
}

var (
	ErrGetTokenId = errors.New("error to get id from token")
	ErrIdDestiny  = errors.New("account destiny id can't be the same account origin id")
)

func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	accountId, ok := middlewares.GetAccountId(r.Context())

	if !ok || accountId == "" {
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	account, err := h.accUse.ListAccountById(accountId)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	var body TransferRequest

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	transaction := transfers.Transfer{
		AccountOriginId:      account.Id,
		AccountDestinationId: body.AccountDestinationId,
		Amount:               body.Amount,
	}

	if string(accountId) == transaction.AccountDestinationId {
		response.SendError(w, ErrIdDestiny, http.StatusBadRequest)
		return
	}

	transfer, err := h.transfUse.CreateTransfer(transaction)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	err = h.transfUse.SaveTransfer(transfer)

	if err != nil {
		log.Printf("Failed to save transfer: %s\n", err.Error())
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	transferResponse := TransferResponse{
		Id:                   transfer.Id,
		AccountOriginId:      transfer.AccountOriginId,
		AccountDestinationId: transfer.AccountDestinationId,
		Amount:               transfer.Amount,
		CreatedAt:            transfer.CreatedAt.Format(response.DateLayout),
	}

	response.SendRequest(w, transferResponse, http.StatusOK)
	log.Printf("sent successful response for transfer %s\n", transfer.Id)
}
