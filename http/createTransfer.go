package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
)

type TransferRequest struct {
	AccountDestinationId string `json:"account_Destination_Id"`
	Amount               int    `json:"amount"`
}

func (s Server) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	accountId, _ := VerifyAuth(w, r)

	account, err := s.app.ListAccountById(string(accountId))

	if err != nil {
		response := Error{Reason: err.Error()}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	var body TransferRequest

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := transfers.Transfer{
		AccountOriginId:      account.Id,
		AccountDestinationId: body.AccountDestinationId,
		Amount:               body.Amount,
	}

	if string(accountId) == transaction.AccountDestinationId {
		response := Error{Reason: "Account destiny id can't be the same account origin id"}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	transfer, err := s.tr.CreateTransfer(transaction)

	if err != nil {
		log.Printf("Failed to do transfer: %s\n", err.Error())
		response := Error{Reason: err.Error()}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	err = s.tr.SaveTransfer(transfer)

	if err != nil {
		log.Printf("Failed to save transfer: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := TransferResponse{
		Id:                   transfer.Id,
		AccountOriginId:      transfer.AccountOriginId,
		AccountDestinationId: transfer.AccountDestinationId,
		Amount:               transfer.Amount,
		CreatedAt:            transfer.CreatedAt.Format(DateLayout),
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent successful response for transfer %s\n", transfer.Id)
}
