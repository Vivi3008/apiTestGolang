package bills

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	accUsecase "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	billUsecase "github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"gotest.tools/v3/assert"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/google/uuid"
)

func TestListBills(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		args               string
		authContextId      middlewares.AuthContextKey
		accountMock        accUsecase.UsecaseMock
		billMock           billUsecase.UsecaseMock
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	personAccount := account.Account{
		Id:        uuid.NewString(),
		Name:      "Teste1",
		Cpf:       "77845100032",
		Secret:    "dafd33255",
		Balance:   250000,
		CreatedAt: time.Now(),
	}

	listBills := []bills.Bill{
		{
			Id:            uuid.NewString(),
			AccountId:     uuid.NewString(),
			Description:   "Academia",
			Value:         6000,
			DueDate:       time.Now(),
			ScheduledDate: time.Now(),
			StatusBill:    bills.Pago,
			CreatedAt:     time.Now(),
		},
		{
			Id:            uuid.NewString(),
			AccountId:     uuid.NewString(),
			Description:   "Internet",
			Value:         15000,
			DueDate:       time.Now(),
			ScheduledDate: time.Now(),
			StatusBill:    bills.Pago,
			CreatedAt:     time.Now(),
		},
	}

	testCases := []TestCase{
		{
			Name:          "Send 200 and list bills succesfull",
			args:          uuid.NewString(),
			authContextId: middlewares.ContextAccountID,
			accountMock: accUsecase.UsecaseMock{
				OnListById: func(accountId string) (account.Account, error) {
					return personAccount, nil
				},
			},
			billMock: billUsecase.UsecaseMock{
				OnListAll: func(id string) ([]bills.Bill, error) {
					return listBills, nil
				},
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want: []BillReqRes{
				{
					Id:            listBills[0].Id,
					AccountId:     listBills[0].AccountId,
					Description:   listBills[0].Description,
					Value:         listBills[0].Value,
					DueDate:       listBills[0].DueDate,
					ScheduledDate: listBills[0].ScheduledDate,
					StatusBill:    listBills[0].StatusBill,
					CreatedAt:     listBills[0].CreatedAt.Format(response.DateLayout),
				},
				{
					Id:            listBills[1].Id,
					AccountId:     listBills[1].AccountId,
					Description:   listBills[1].Description,
					Value:         listBills[1].Value,
					DueDate:       listBills[1].DueDate,
					ScheduledDate: listBills[1].ScheduledDate,
					StatusBill:    listBills[1].StatusBill,
					CreatedAt:     listBills[1].CreatedAt.Format(response.DateLayout),
				},
			},
		},
		{
			Name:          "Sent error 500 if error in usecase",
			args:          uuid.NewString(),
			authContextId: middlewares.ContextAccountID,
			accountMock:   accUsecase.UsecaseMock{},
			billMock: billUsecase.UsecaseMock{
				OnListAll: func(id string) ([]bills.Bill, error) {
					return []bills.Bill{}, fmt.Errorf("account id doesn't exist")
				},
			},
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "account id doesn't exist",
			},
		},
		{
			Name:          "Sent error 401 if invalid auth",
			args:          uuid.NewString(),
			authContextId: "id",
			accountMock:   accUsecase.UsecaseMock{},
			billMock: billUsecase.UsecaseMock{
				OnListAll: func(id string) ([]bills.Bill, error) {
					return []bills.Bill{}, nil
				},
			},
			wantHttpStatusCode: 401,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: ErrGetTokenId.Error(),
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{blUse: tt.billMock, accUse: tt.accountMock}
			request := httptest.NewRequest(http.MethodGet, "/bills", nil)

			ctx := context.WithValue(request.Context(), tt.authContextId, tt.args)
			resp := httptest.NewRecorder()

			wantBody, _ := json.Marshal(tt.want)

			http.HandlerFunc(handler.ListBills).ServeHTTP(resp, request.WithContext(ctx))

			assert.Equal(t, string(wantBody), strings.TrimSpace(resp.Body.String()))
			assert.Equal(t, tt.wantHeader, resp.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, resp.Code)
		})
	}
}
