package transfers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	lg "github.com/Vivi3008/apiTestGolang/gateways/http/logging"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
)

type TransferRequest struct {
	AccountDestinationId string `json:"account_Destination_Id"`
	Amount               int    `json:"amount"`
}

var (
	ErrGetTokenId             = errors.New("error to get id from token")
	ErrIdDestiny              = errors.New("account destiny id can't be the same account origin id")
	ErrInvalidTransferPayload = errors.New("invalid transfer payload")
)

func (h Handler) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.transfers.CreateTransfer"
	accountId, ok := middlewares.GetAccountId(r.Context())

	log := lg.FromContext(r.Context(), operation)

	if !ok || accountId == "" {
		log.Error("Error to get token from id")
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	var body TransferRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		log.WithError(err).Error("Error in request body")
		response.SendError(w, ErrInvalidTransferPayload, http.StatusBadRequest)
		return
	}

	transaction := transfers.Transfer{
		AccountOriginId:      accountId,
		AccountDestinationId: body.AccountDestinationId,
		Amount:               body.Amount,
	}

	log.Info("Starting create transfer")
	transfer, err := h.transfUse.CreateTransfer(r.Context(), transaction)

	if err != nil {
		log.WithError(err).Error("Error to create transfer")
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	log.Info("Starting save transfer")
	err = h.transfUse.SaveTransfer(r.Context(), transfer)

	if err != nil {
		log.WithError(err).Error("Error to save transfer")
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

	response.Send(w, transferResponse, http.StatusOK)
	log.WithField("transferId", transferResponse.Id).Info("Transfer created successfully")
}
