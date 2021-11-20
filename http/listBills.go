package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain"
)

func (s Server) ListBills(w http.ResponseWriter, r *http.Request) {
	accountId, _ := VerifyAuth(w, r)

	list, err := s.bl.ListAllBills(domain.AccountId(accountId))

	if err != nil {
		log.Printf("Failed to do list bills: %s\n", err.Error())
		return
	}

	response := make([]BillReqRes, len(list))

	for i, bill := range list {
		response[i].AccountId = bill.AccountId
		response[i].Description = bill.Description
		response[i].Value = bill.Value
		response[i].DueDate = bill.DueDate
		response[i].ScheduledDate = bill.ScheduledDate
		response[i].StatusBill = bill.StatusBill
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("Sent all bills from Id %s", accountId)
	log.Printf("Sent all bills. Total: %d", len(response))
}
