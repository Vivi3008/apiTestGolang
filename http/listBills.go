package http

import (
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/http/response"
)

func (s Server) ListBills(w http.ResponseWriter, r *http.Request) {
	accountId, ok := middlewares.GetAccountId(r.Context())

	if !ok || accountId == "" {
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	list, err := s.bl.ListAllBills(string(accountId))

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	payments := make([]BillReqRes, len(list))

	for i, bill := range list {
		payments[i].AccountId = bill.AccountId
		payments[i].Description = bill.Description
		payments[i].Value = bill.Value
		payments[i].DueDate = bill.DueDate
		payments[i].ScheduledDate = bill.ScheduledDate
		payments[i].StatusBill = bill.StatusBill
	}

	response.SendRequest(w, payments, http.StatusOK)
	log.Printf("Sent all bills from Id %s", accountId)
	log.Printf("Sent all bills. Total: %d", len(payments))
}
