package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain"
)

type TransferRequest struct {
	AccountDestinationId string  `json:"account_Destination_Id"`
	Amount               float64 `json:"amount"`
}

func (s Server) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	accountId, _ := VerifyAuth(w, r)

	var body TransferRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := domain.Transfer{
		AccountOriginId:      string(accountId),
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

	transfer, err := s.app.CreateTransfer(transaction)

	if err != nil {
		log.Printf("Failed to do transfer: %s\n", err.Error())
		response := Error{Reason: err.Error()}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	saveTransfer, err := s.tr.SaveTransfer(transfer)

	if err != nil {
		log.Printf("Failed to save transfer: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := TransferResponse{
		Id:                   saveTransfer.Id,
		AccountOriginId:      saveTransfer.AccountOriginId,
		AccountDestinationId: saveTransfer.AccountDestinationId,
		Amount:               saveTransfer.Amount,
		CreatedAt:            saveTransfer.CreatedAt,
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent successful response for transfer %s\n", transfer.Id)
}
