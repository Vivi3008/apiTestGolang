package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

type BillReqRes struct {
	AccountId     string       `json:"account_id"`
	Description   string       `json:"description"`
	Value         int          `json:"value"`
	DueDate       time.Time    `json:"due_date"`
	ScheduledDate time.Time    `json:"scheduled_date"`
	StatusBill    bills.Status `json:"status"`
}

func (s Server) CreateBill(w http.ResponseWriter, r *http.Request) {
	accountId, _ := VerifyAuth(w, r)

	var body BillReqRes

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	bill := bills.Bill{
		AccountId:   string(accountId),
		Description: body.Description,
		Value:       body.Value,
		DueDate:     body.DueDate,
	}

	billOk, err := s.bl.CreateBill(bill)

	if err != nil {
		log.Printf("Failed to pay bill: %s\n", err.Error())
		response := Error{Reason: err.Error()}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	newBill := bills.Bill{
		Id:          billOk.Id,
		AccountId:   billOk.AccountId,
		Description: billOk.Description,
		Value:       billOk.Value,
		DueDate:     billOk.DueDate,
		StatusBill:  bills.Pago,
	}

	err = s.bl.SaveBill(newBill)

	if err != nil {
		log.Printf("Failed to save bill: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	billResonse := BillReqRes{
		AccountId:     newBill.AccountId,
		Description:   newBill.Description,
		Value:         newBill.Value,
		DueDate:       newBill.DueDate,
		ScheduledDate: newBill.ScheduledDate,
		StatusBill:    newBill.StatusBill,
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(billResonse)
	log.Printf("sent successful response for transfer %s\n", newBill.Id)
}
