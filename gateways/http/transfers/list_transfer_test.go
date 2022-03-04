package transfers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	accEntitie "github.com/Vivi3008/apiTestGolang/domain/entities/account"
	entities "github.com/Vivi3008/apiTestGolang/domain/entities/transfers"

	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/google/uuid"
	"gotest.tools/v3/assert"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
)

func TestListTransfers(t *testing.T) {
	t.Parallel()

	accountOrigin := uuid.NewString()

	transferTest := []entities.Transfer{
		{
			Id:                   uuid.NewString(),
			AccountOriginId:      accountOrigin,
			AccountDestinationId: uuid.NewString(),
			Amount:               60000,
			CreatedAt:            time.Now(),
		},
		{
			Id:                   uuid.NewString(),
			AccountOriginId:      accountOrigin,
			AccountDestinationId: uuid.NewString(),
			Amount:               70000,
			CreatedAt:            time.Now(),
		},
	}

	personAccount := accEntitie.Account{
		Id:        uuid.NewString(),
		Name:      "Teste1",
		Cpf:       "77845100032",
		Secret:    "dafd33255",
		Balance:   250000,
		CreatedAt: time.Now(),
	}

	type TestCase struct {
		Name               string
		args               string
		accountMock        account.UsecaseMock
		transferMock       transfers.UsecaseMock
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	testCases := []TestCase{
		{
			Name: "Should sent list transfers successful and status code 200",
			accountMock: account.UsecaseMock{
				OnListById: func(accountId string) (accEntitie.Account, error) {
					return personAccount, nil
				},
			},
			transferMock: transfers.UsecaseMock{
				OnListAll: func(id string) ([]entities.Transfer, error) {
					return transferTest, nil
				},
			},
			args:               accountOrigin,
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want: []TransferResponse{
				{
					Id:                   transferTest[0].Id,
					AccountOriginId:      transferTest[0].AccountOriginId,
					AccountDestinationId: transferTest[0].AccountDestinationId,
					Amount:               transferTest[0].Amount,
					CreatedAt:            transferTest[0].CreatedAt.Format(response.DateLayout),
				},
				{
					Id:                   transferTest[1].Id,
					AccountOriginId:      transferTest[1].AccountOriginId,
					AccountDestinationId: transferTest[1].AccountDestinationId,
					Amount:               transferTest[1].Amount,
					CreatedAt:            transferTest[1].CreatedAt.Format(response.DateLayout),
				},
			},
		},
		{
			Name: "Sent 500 if error in account usecase",
			accountMock: account.UsecaseMock{
				OnListById: func(accountId string) (accEntitie.Account, error) {
					return accEntitie.Account{}, fmt.Errorf("account id doesn't exist")
				},
			},
			transferMock: transfers.UsecaseMock{
				OnListAll: func(id string) ([]entities.Transfer, error) {
					return transferTest, nil
				},
			},
			args:               accountOrigin,
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "account id doesn't exist",
			},
		},
		{
			Name: "Sent 500 if error in transfer usecase",
			accountMock: account.UsecaseMock{
				OnListById: func(accountId string) (accEntitie.Account, error) {
					return personAccount, nil
				},
			},
			transferMock: transfers.UsecaseMock{
				OnListAll: func(id string) ([]entities.Transfer, error) {
					return []entities.Transfer{}, fmt.Errorf("error in database")
				},
			},
			args:               accountOrigin,
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "error in database",
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{accUse: tt.accountMock, transfUse: tt.transferMock}
			accountContextId := middlewares.ContextAccountID

			req := httptest.NewRequest(http.MethodGet, "/transfers", nil)
			ctx := context.WithValue(req.Context(), accountContextId, tt.args)

			wantBody, _ := json.Marshal(tt.want)

			resp := httptest.NewRecorder()

			http.HandlerFunc(handler.ListTransfer).ServeHTTP(resp, req.WithContext(ctx))

			assert.Equal(t, string(wantBody), strings.TrimSpace(resp.Body.String()))
			assert.Equal(t, tt.wantHeader, resp.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, resp.Code)
		})
	}
}
