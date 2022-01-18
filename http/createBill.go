package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/http/response"
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
	accountId, ok := middlewares.GetAccountId(r.Context())

	if !ok || accountId == "" {
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	var body BillReqRes

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
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
		response.SendError(w, err, http.StatusBadRequest)
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
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	billResponse := BillReqRes{
		AccountId:     newBill.AccountId,
		Description:   newBill.Description,
		Value:         newBill.Value,
		DueDate:       newBill.DueDate,
		ScheduledDate: newBill.ScheduledDate,
		StatusBill:    newBill.StatusBill,
	}

	response.SendRequest(w, billResponse, http.StatusOK)
	log.Printf("sent successful response for transfer %s\n", newBill.Id)
}
