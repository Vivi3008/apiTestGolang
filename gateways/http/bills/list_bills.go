package bills

import (
	"net/http"

	lg "github.com/Vivi3008/apiTestGolang/gateways/http/logging"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/sirupsen/logrus"
)

func (h Handler) ListBills(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.bills.ListBills"
	accountId, ok := middlewares.GetAccountId(r.Context())

	log := lg.FromContext(r.Context(), operation)

	if !ok || accountId == "" {
		log.Error("Error to get id from token")
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	log.Info("Starting list bills")
	list, err := h.blUse.ListBills(r.Context(), accountId)

	if err != nil {
		log.WithError(err).Error("Error to list bills")
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	payments := make([]BillReqRes, len(list))

	for i, bill := range list {
		payments[i].Id = bill.Id
		payments[i].AccountId = bill.AccountId
		payments[i].Description = bill.Description
		payments[i].Value = bill.Value
		payments[i].DueDate = bill.DueDate
		payments[i].ScheduledDate = bill.ScheduledDate
		payments[i].StatusBill = bill.StatusBill
		payments[i].CreatedAt = bill.CreatedAt.Format(response.DateLayout)
	}

	response.Send(w, payments, http.StatusOK)
	log.WithFields(logrus.Fields{
		"account_id": accountId,
		"total":      len(payments),
	}).Info("Bills listed succesfully")
}
