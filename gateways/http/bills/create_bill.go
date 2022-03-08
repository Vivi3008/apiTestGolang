package bills

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	lg "github.com/Vivi3008/apiTestGolang/gateways/http/logging"
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
	const operation = "handler.bills.CreateBill"
	accountId, ok := middlewares.GetAccountId(r.Context())

	log := lg.FromContext(r.Context(), operation)

	if !ok || accountId == "" {
		log.Error("Error to get account id from token")
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	var body BillReqRes

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		log.WithError(err).Error("Error in request body")
		response.SendError(w, ErrInvalidBillPayload, http.StatusBadRequest)
		return
	}

	bill := bills.Bill{
		AccountId:   string(accountId),
		Description: body.Description,
		Value:       body.Value,
		DueDate:     body.DueDate,
	}

	log.Info("Starting create a bill")
	billOk, err := h.blUse.CreateBill(r.Context(), bill)

	if err != nil {
		log.WithError(err).Error("Error to create bill")
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	billOk.StatusBill = bills.Pago

	log.Info("Saving bill created")
	err = h.blUse.SaveBill(r.Context(), billOk)

	if err != nil {
		log.WithError(err).Error("Error to save bill")
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
	log.WithField("billId", billOk.Id).Info("Create bill successfull")
}
