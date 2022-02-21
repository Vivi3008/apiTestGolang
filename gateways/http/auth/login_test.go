package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestLogin(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		accountMock        account.AccountMock
		args               interface{}
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	id := uuid.NewString()
	token, err := commom.CreateToken(id)

	if err != nil {
		t.Errorf("Error to create token: %s", err)
	}

	testCases := []TestCase{
		{
			Name: "Should log in successfull return 200",
			accountMock: account.AccountMock{
				OnLogin: func(u account.Login) (string, error) {
					return id, nil
				},
			},
			args: LoginRequest{
				Cpf:    "12345678910",
				Secret: "123456",
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want: TokenString{
				Token: token,
			},
		},
		{
			Name: "Return 400 if credentials is invalid, error in usecase",
			accountMock: account.AccountMock{
				OnLogin: func(u account.Login) (string, error) {
					return "", fmt.Errorf("invalid password")
				},
			},
			args: LoginRequest{
				Cpf:    "12345678910",
				Secret: "123456",
			},
			wantHttpStatusCode: 401,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "invalid password",
			},
		},
		{
			Name: "Return 400 if invalid body in request",
			accountMock: account.AccountMock{
				OnLogin: func(u account.Login) (string, error) {
					return uuid.NewString(), nil
				},
			},
			args:               "not valid body",
			wantHttpStatusCode: 400,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: ErrInvalidLoginPayload.Error(),
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{accUse: tt.accountMock}

			reqBody, _ := json.Marshal(tt.args)
			req := bytes.NewReader(reqBody)

			request := httptest.NewRequest(http.MethodPost, "/login", req)
			resp := httptest.NewRecorder()

			wantBody, _ := json.Marshal(tt.want)
			http.HandlerFunc(handler.Login).ServeHTTP(resp, request)

			assert.Equal(t, string(wantBody), strings.TrimSpace(resp.Body.String()))
			assert.Equal(t, tt.wantHeader, resp.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, resp.Code)
		})
	}
}
