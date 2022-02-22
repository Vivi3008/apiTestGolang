package accounts

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	usecase "github.com/Vivi3008/apiTestGolang/domain/usecases/account"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

var ListAccounts = []account.Account{
	{
		Id:        uuid.NewString(),
		Name:      "Teste1",
		Cpf:       "77845100032",
		Secret:    "dafd33255",
		Balance:   250000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste2",
		Cpf:       "55985633301",
		Secret:    "4f5ds4af54",
		Balance:   260000,
		CreatedAt: time.Now(),
	},
	{
		Id:        uuid.NewString(),
		Name:      "Teste3",
		Cpf:       "85665232145",
		Secret:    "fadsfdsaf",
		Balance:   360000,
		CreatedAt: time.Now(),
	},
}

func TestListAccounts(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		accountMock        usecase.UsecaseMock
		wantHttpStatusCode int
		wantHeader         string
		wantBody           interface{}
	}

	testCases := []TestCase{
		{
			Name: "Should list accounts succesfull",
			accountMock: usecase.UsecaseMock{
				OnListAll: func() ([]account.Account, error) {
					return ListAccounts, nil
				},
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			wantBody: []AccountResponse{
				{
					Id:        ListAccounts[0].Id,
					Name:      ListAccounts[0].Name,
					Cpf:       ListAccounts[0].Cpf,
					Balance:   ListAccounts[0].Balance,
					CreatedAt: ListAccounts[0].CreatedAt.Format(response.DateLayout),
				},
				{
					Id:        ListAccounts[1].Id,
					Name:      ListAccounts[1].Name,
					Cpf:       ListAccounts[1].Cpf,
					Balance:   ListAccounts[1].Balance,
					CreatedAt: ListAccounts[1].CreatedAt.Format(response.DateLayout),
				},
				{
					Id:        ListAccounts[2].Id,
					Name:      ListAccounts[2].Name,
					Cpf:       ListAccounts[2].Cpf,
					Balance:   ListAccounts[2].Balance,
					CreatedAt: ListAccounts[2].CreatedAt.Format(response.DateLayout),
				},
			},
		},
		{
			Name: "Fail in list accounts whit internal server error",
			accountMock: usecase.UsecaseMock{
				OnListAll: func() ([]account.Account, error) {
					return []account.Account{}, fmt.Errorf("error in connection database")
				},
			},
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			wantBody: response.Error{
				Reason: "error in connection database",
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{acc: tt.accountMock}
			request := httptest.NewRequest(http.MethodGet, "/accounts", nil)
			response := httptest.NewRecorder()

			http.HandlerFunc(handler.ListAll).ServeHTTP(response, request)

			wantBody, _ := json.Marshal(tt.wantBody)

			assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
			assert.Equal(t, tt.wantHeader, response.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, response.Code)
		})
	}
}
