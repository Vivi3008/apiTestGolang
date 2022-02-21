package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"gotest.tools/v3/assert"
)

func TestGetBalance(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		accountMock        account.AccountMock
		args               interface{}
		wantBody           interface{}
		wantHttpStatusCode int
		wantHeader         string
	}

	testCases := []TestCase{
		{
			Name: "Should sent balance succesfull",
			accountMock: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return ListAccounts[0], nil
				},
			},
			args: AccountIdRequest{Id: ListAccounts[0].Id},
			wantBody: BalanceAccountResponse{
				Balance: ListAccounts[0].Balance,
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
		},
		{
			Name: "Return 500 if error in usecase",
			accountMock: account.AccountMock{
				OnListById: func(accountId string) (account.Account, error) {
					return account.Account{}, fmt.Errorf("id doesn't exist")
				},
			},
			args: AccountIdRequest{Id: ListAccounts[0].Id},
			wantBody: response.Error{
				Reason: "id doesn't exist",
			},
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{acc: tt.accountMock}
			path := fmt.Sprintf("/accounts/{%s}/balance", tt.args)

			request := httptest.NewRequest(http.MethodGet, path, nil)
			response := httptest.NewRecorder()

			http.HandlerFunc(handler.GetBalance).ServeHTTP(response, request)

			wantBody, _ := json.Marshal(tt.wantBody)

			assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
			assert.Equal(t, tt.wantHeader, response.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, response.Code)
		})
	}
}
