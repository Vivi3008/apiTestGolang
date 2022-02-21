package bills

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestCreateBill(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		bodyArgs           interface{}
		accountMock        account.AccountMock
		billMock           bills.BillMock
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	billTest := bills.Bill{
		AccountId:   uuid.NewString(),
		Description: "Academia",
		Value:       10000,
		DueDate:     time.Now(),
	}
	bill, _ := bills.NewBill(billTest)

	customer := account.Account{
		Name:    "Testando",
		Cpf:     "13265478951",
		Secret:  "123456",
		Balance: 5000000,
	}

	personAccount, _ := account.NewAccount(customer)

	testCases := []TestCase{
		{
			Name: "Should create a bill succesfull and return 200",
			accountMock: account.AccountMock{
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return personAccount, nil
				},
			},
			billMock: bills.BillMock{
				OnCreate: func(b bills.Bill) (bills.Bill, error) {
					return bill, nil
				},
				OnStore: func(b bills.Bill) error {
					return nil
				},
			},
			bodyArgs: BillReqRes{
				AccountId:   bill.AccountId,
				Description: bill.Description,
				Value:       bill.Value,
				DueDate:     bill.DueDate,
			},
			want: BillReqRes{
				Id:            bill.Id,
				AccountId:     bill.AccountId,
				Description:   bill.Description,
				Value:         bill.Value,
				DueDate:       bill.DueDate,
				ScheduledDate: bill.ScheduledDate,
				StatusBill:    bills.Pago,
				CreatedAt:     bill.CreatedAt.Format(response.DateLayout),
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
		},
		{
			Name: "Return 500 if error in database",
			accountMock: account.AccountMock{
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return personAccount, nil
				},
			},
			billMock: bills.BillMock{
				OnCreate: func(b bills.Bill) (bills.Bill, error) {
					return bill, nil
				},
				OnStore: func(b bills.Bill) error {
					return fmt.Errorf("error to save in database")
				},
			},
			bodyArgs: BillReqRes{
				AccountId:   bill.AccountId,
				Description: bill.Description,
				Value:       bill.Value,
				DueDate:     bill.DueDate,
			},
			want: response.Error{
				Reason: "error to save in database",
			},
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
		},
		{
			Name: "Return 400 if error in usecase",
			accountMock: account.AccountMock{
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return account.Account{}, fmt.Errorf("insuficient limit")
				},
			},
			billMock: bills.BillMock{
				OnCreate: func(b bills.Bill) (bills.Bill, error) {
					return bills.Bill{}, fmt.Errorf("insuficient limit")
				},
				OnStore: func(b bills.Bill) error {
					return nil
				},
			},
			bodyArgs: BillReqRes{
				AccountId:   bill.AccountId,
				Description: bill.Description,
				Value:       bill.Value,
				DueDate:     bill.DueDate,
			},
			want: response.Error{
				Reason: "insuficient limit",
			},
			wantHttpStatusCode: 400,
			wantHeader:         response.JSONContentType,
		},
		{
			Name: "Return 400 if invalid body in request",
			accountMock: account.AccountMock{
				OnUpdade: func(balance int, id string) (account.Account, error) {
					return personAccount, nil
				},
			},
			billMock: bills.BillMock{
				OnCreate: func(b bills.Bill) (bills.Bill, error) {
					return bill, nil
				},
				OnStore: func(b bills.Bill) error {
					return nil
				},
			},
			bodyArgs: "qualquer coisa",
			want: response.Error{
				Reason: ErrInvalidBillPayload.Error(),
			},
			wantHttpStatusCode: 400,
			wantHeader:         response.JSONContentType,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{blUse: tt.billMock, accUse: tt.accountMock}
			reqBody, _ := json.Marshal(tt.bodyArgs)
			req := bytes.NewReader(reqBody)

			authContextId := middlewares.ContextAccountID

			wantBody, _ := json.Marshal(tt.want)
			request := httptest.NewRequest(http.MethodPost, "/bills", req)
			resp := httptest.NewRecorder()
			ctx := context.WithValue(request.Context(), authContextId, personAccount.Id)

			http.HandlerFunc(handler.CreateBill).ServeHTTP(resp, request.WithContext(ctx))

			assert.Equal(t, string(wantBody), strings.TrimSpace(resp.Body.String()))
			assert.Equal(t, tt.wantHeader, resp.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, resp.Code)
		})
	}
}
