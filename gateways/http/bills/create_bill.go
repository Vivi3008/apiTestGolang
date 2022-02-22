package bills

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
)

type BillReqRes struct {
	Id            string       `json:"id"`
	AccountId     string       `json:"account_id"`
	Description   string       `json:"description"`
	Value         int          `json:"value"`
	DueDate       time.Time    `json:"due_date"`
	ScheduledDate time.Time    `json:"scheduled_date"`
	StatusBill    bills.Status `json:"status"`
	CreatedAt     string       `json:"created_at"`
}

var (
	ErrGetTokenId         = errors.New("error to get id from token")
	ErrInvalidBillPayload = errors.New("invalid bill payload")
)

func (h Handler) CreateBill(w http.ResponseWriter, r *http.Request) {
	accountId, ok := middlewares.GetAccountId(r.Context())

	if !ok || accountId == "" {
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	var body BillReqRes

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.SendError(w, ErrInvalidBillPayload, http.StatusBadRequest)
		return
	}

	bill := bills.Bill{
		AccountId:   string(accountId),
		Description: body.Description,
		Value:       body.Value,
		DueDate:     body.DueDate,
	}

	billOk, err := h.blUse.CreateBill(r.Context(), bill)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	billOk.StatusBill = bills.Pago

	err = h.blUse.SaveBill(r.Context(), billOk)

	if err != nil {
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	billResponse := BillReqRes{
		Id:            billOk.Id,
		AccountId:     billOk.AccountId,
		Description:   billOk.Description,
		Value:         billOk.Value,
		DueDate:       billOk.DueDate,
		ScheduledDate: billOk.ScheduledDate,
		StatusBill:    billOk.StatusBill,
		CreatedAt:     billOk.CreatedAt.Format(response.DateLayout),
	}

	response.Send(w, billResponse, http.StatusOK)
	log.Printf("sent successful response for bill %s\n", billOk.Id)
}
