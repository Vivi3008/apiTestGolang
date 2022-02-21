package transfers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	entities "github.com/Vivi3008/apiTestGolang/domain/entities/transfers"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestCreateTransfer(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		args               interface{}
		accountMock        account.UsecaseMock
		transferMock       transfers.UsecaseMock
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	transReq := TransferRequest{
		AccountDestinationId: uuid.NewString(),
		Amount:               60000,
	}

	newTransfer, err := entities.NewTransfer(entities.Transfer{
		AccountOriginId:      uuid.NewString(),
		AccountDestinationId: transReq.AccountDestinationId,
		Amount:               transReq.Amount,
	})

	if err != nil {
		t.Errorf("Error %s", err)
	}

	testCases := []TestCase{
		{
			Name:        "Should create transfers succesfull and sent a transfer with status code 200",
			args:        transReq,
			accountMock: account.UsecaseMock{},
			transferMock: transfers.UsecaseMock{
				OnCreate: func(trans entities.Transfer) (entities.Transfer, error) {
					return newTransfer, nil
				},
				OnSave: func(trans entities.Transfer) error {
					return nil
				},
			},
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want: TransferResponse{
				Id:                   newTransfer.Id,
				AccountOriginId:      newTransfer.AccountOriginId,
				AccountDestinationId: newTransfer.AccountDestinationId,
				Amount:               newTransfer.Amount,
				CreatedAt:            newTransfer.CreatedAt.Format(response.DateLayout),
			},
		},
		{
			Name:        "Sent error 500 if error in usecase",
			args:        transReq,
			accountMock: account.UsecaseMock{},
			transferMock: transfers.UsecaseMock{
				OnCreate: func(trans entities.Transfer) (entities.Transfer, error) {
					return newTransfer, nil
				},
				OnSave: func(trans entities.Transfer) error {
					return fmt.Errorf("error to save in database")
				},
			},
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "error to save in database",
			},
		},
		{
			Name:        "Sent error 500 if error in create usecase",
			args:        transReq,
			accountMock: account.UsecaseMock{},
			transferMock: transfers.UsecaseMock{
				OnCreate: func(trans entities.Transfer) (entities.Transfer, error) {
					return entities.Transfer{}, fmt.Errorf("origin id is equal destiny id")
				},
				OnSave: func(trans entities.Transfer) error {
					return nil
				},
			},
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "origin id is equal destiny id",
			},
		},
		{
			Name:        "Sent error 400 if client sent invalid payload",
			args:        "invalid body",
			accountMock: account.UsecaseMock{},
			transferMock: transfers.UsecaseMock{
				OnCreate: func(trans entities.Transfer) (entities.Transfer, error) {
					return newTransfer, nil
				},
				OnSave: func(trans entities.Transfer) error {
					return nil
				},
			},
			wantHttpStatusCode: 400,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: ErrInvalidTransferPayload.Error(),
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{accUse: tc.accountMock, transfUse: tt.transferMock}

			authContextId := middlewares.ContextAccountID

			reqBody, _ := json.Marshal(tt.args)
			req := bytes.NewReader(reqBody)
			request := httptest.NewRequest(http.MethodPost, "/transfers", req)

			ctx := context.WithValue(request.Context(), authContextId, newTransfer.AccountOriginId)

			wantBody, _ := json.Marshal(tt.want)
			resp := httptest.NewRecorder()

			http.HandlerFunc(handler.CreateTransfer).ServeHTTP(resp, request.WithContext(ctx))

			assert.Equal(t, string(wantBody), strings.TrimSpace(resp.Body.String()))
			assert.Equal(t, tt.wantHeader, resp.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, resp.Code)
		})
	}
}
