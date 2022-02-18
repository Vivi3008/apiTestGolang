package activities

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"

	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"

	"github.com/google/uuid"
	"gotest.tools/v3/assert"
)

func TestListAcitivities(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name               string
		args               string
		authContextId      middlewares.AuthContextKey
		activitiesMock     activities.AcitivitiesMock
		wantHttpStatusCode int
		wantHeader         string
		want               interface{}
	}

	listActivities := []activities.AccountActivity{
		{
			Type:   activities.Bill,
			Amount: 1330000,
			CreatedAt: time.Date(2022, time.February,
				13, 21, 34, 01, 0, time.UTC),
			Details: activities.DescriptionPayment{
				Description: "Academia",
				Status:      bills.Pago,
			},
		},
		{
			Type:   activities.Bill,
			Amount: 500000,
			CreatedAt: time.Date(2022, time.February,
				13, 21, 34, 01, 0, time.UTC),
			Details: activities.DescriptionPayment{
				Description: "Iptu",
				Status:      bills.Agendado,
			},
		},
		{
			Type:   activities.Bill,
			Amount: 462222,
			CreatedAt: time.Date(2022, time.February,
				13, 21, 34, 01, 0, time.UTC),
			Details: activities.DescriptionPayment{
				Description: "Internet",
				Status:      bills.Negado,
			},
		},
		{
			Type:   activities.Transfer,
			Amount: 500000,
			CreatedAt: time.Date(2022, time.February,
				13, 21, 34, 01, 0, time.UTC),
			Details: activities.DestinyAccount{
				Name:                 "Cindy",
				AccountDestinationId: uuid.NewString(),
			},
		},
		{
			Type:   activities.Transfer,
			Amount: 600000,
			CreatedAt: time.Date(2022, time.February,
				13, 21, 34, 01, 0, time.UTC),
			Details: activities.DestinyAccount{
				Name:                 "Branca",
				AccountDestinationId: uuid.NewString(),
			},
		},
	}

	testCases := []TestCase{
		{
			Name: "Should sent activities successfull",
			activitiesMock: activities.AcitivitiesMock{
				OnListActivities: func(accountId string) ([]activities.AccountActivity, error) {
					return listActivities, nil
				},
			},
			args:               uuid.NewString(),
			authContextId:      middlewares.ContextAccountID,
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want: []ActivitiesResponse{
				{
					Type:      listActivities[0].Type,
					Amount:    listActivities[0].Amount,
					CreatedAt: listActivities[0].CreatedAt,
					Details:   listActivities[0].Details,
				},
				{
					Type:      listActivities[1].Type,
					Amount:    listActivities[1].Amount,
					CreatedAt: listActivities[1].CreatedAt,
					Details:   listActivities[1].Details,
				},
				{
					Type:      listActivities[2].Type,
					Amount:    listActivities[2].Amount,
					CreatedAt: listActivities[2].CreatedAt,
					Details:   listActivities[2].Details,
				},
				{
					Type:      listActivities[3].Type,
					Amount:    listActivities[3].Amount,
					CreatedAt: listActivities[3].CreatedAt,
					Details:   listActivities[3].Details,
				},
				{
					Type:      listActivities[4].Type,
					Amount:    listActivities[4].Amount,
					CreatedAt: listActivities[4].CreatedAt,
					Details:   listActivities[4].Details,
				},
			},
		},
		{
			Name: "Fail if token is invalid",
			activitiesMock: activities.AcitivitiesMock{
				OnListActivities: func(accountId string) ([]activities.AccountActivity, error) {
					return listActivities, nil
				},
			},
			args:               "131312",
			authContextId:      "id",
			wantHttpStatusCode: 401,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: ErrGetTokenId.Error(),
			},
		},
		{
			Name: "Sent list activities empty if id doesn't exist",
			activitiesMock: activities.AcitivitiesMock{
				OnListActivities: func(accountId string) ([]activities.AccountActivity, error) {
					return []activities.AccountActivity{}, nil
				},
			},
			args:               "131312",
			authContextId:      middlewares.ContextAccountID,
			wantHttpStatusCode: 200,
			wantHeader:         response.JSONContentType,
			want:               []activities.AccountActivity{},
		},
		{
			Name: "Sent internal server error if unknown error",
			activitiesMock: activities.AcitivitiesMock{
				OnListActivities: func(accountId string) ([]activities.AccountActivity, error) {
					return []activities.AccountActivity{}, fmt.Errorf("error in database connection")
				},
			},
			args:               uuid.NewString(),
			authContextId:      middlewares.ContextAccountID,
			wantHttpStatusCode: 500,
			wantHeader:         response.JSONContentType,
			want: response.Error{
				Reason: "error in database connection",
			},
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := &Handler{actUse: tt.activitiesMock}
			request := httptest.NewRequest(http.MethodGet, "/activity", nil)

			ctx := context.WithValue(request.Context(), tt.authContextId, tt.args)
			res := httptest.NewRecorder()

			wantBody, _ := json.Marshal(tt.want)

			http.HandlerFunc(handler.ListActivity).ServeHTTP(res, request.WithContext(ctx))

			assert.Equal(t, string(wantBody), strings.TrimSpace(res.Body.String()))
			assert.Equal(t, tt.wantHeader, res.Header().Get("Content-Type"))
			assert.Equal(t, tt.wantHttpStatusCode, res.Code)
		})
	}
}
