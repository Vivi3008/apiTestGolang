package transfers

import (
	"net/http"

	lg "github.com/Vivi3008/apiTestGolang/gateways/http/logging"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/sirupsen/logrus"
)

type TransferResponse struct {
	Id                   string `json:"id"`
	AccountOriginId      string `json:"account_origin_id"`
	AccountDestinationId string `json:"account_destination_id"`
	Amount               int    `json:"amount"`
	CreatedAt            string `json:"createdAt"`
}

func (h Handler) ListTransfer(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.transfers.ListTransfer"
	accountId, ok := middlewares.GetAccountId(r.Context())

	log := lg.FromContext(r.Context(), operation)

	if !ok || accountId == "" {
		log.Error("Error to get token from id")
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	log.Info("Starting to list account")
	account, err := h.accUse.ListAccountById(r.Context(), accountId)

	if err != nil {
		log.WithError(err).Error("Error to list account")
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	log.Info("Starting to list transfer")
	list, err := h.transfUse.ListTransfer(r.Context(), account.Id)

	if err != nil {
		log.WithError(err).Info("Error to list transfer")
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

	response.Send(w, transfers, http.StatusOK)
	log.WithFields(logrus.Fields{
		"account_id": accountId,
		"total":      len(transfers),
	}).Info("List transfers successfully")
}
