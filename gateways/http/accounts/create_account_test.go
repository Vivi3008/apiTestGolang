package accounts

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"gotest.tools/v3/assert"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		bodyArgs           interface{}
		accountMock        account.AccountMock
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	customer := account.Account{
		Name:    "Testando",
		Cpf:     "13265478951",
		Secret:  "123456",
		Balance: 5000000,
	}

	personAccount, _ := account.NewAccount(customer)

	testCases := []TestCase{
		{
			Name: "Should create account successfull send 200",
			accountMock: account.AccountMock{
				OnCreate: func(acc account.Account) (account.Account, error) {
					return personAccount, nil
				},
				OnStoreAccount: func(account account.Account) error {
					return nil
				},
			},
			bodyArgs: AccountRequest{
				Name:    customer.Name,
				Cpf:     customer.Cpf,
				Secret:  customer.Secret,
				Balance: customer.Balance,
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want: AccountResponse{
				Id:        personAccount.Id,
				Name:      personAccount.Name,
				Cpf:       personAccount.Cpf,
				Balance:   personAccount.Balance,
				CreatedAt: personAccount.CreatedAt.Format(response.DateLayout),
			},
		},
		{
			Name: "Fail if body is invalid and return 400",
			accountMock: account.AccountMock{
				OnCreate: func(acc account.Account) (account.Account, error) {
					return personAccount, nil
				},
				OnStoreAccount: func(account account.Account) error {
					return nil
				},
			},
			bodyArgs:           "invalid body",
			wantHttpStatusCode: 400,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: ErrInvalidPayloadAccount.Error(),
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{
				acc: tt.accountMock,
			}
			reqBody, _ := json.Marshal(tt.bodyArgs)
			req := bytes.NewReader(reqBody)
			request := httptest.NewRequest(http.MethodPost, "/accounts", req)

			wantBody, _ := json.Marshal(tt.want)
			response := httptest.NewRecorder()

			http.HandlerFunc(handler.CreateAccount).ServeHTTP(response, request)

			assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
			assert.Equal(t, tt.wantHeader, response.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, response.Code)
		})
	}
}
