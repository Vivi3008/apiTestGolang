package activities

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
	"github.com/google/uuid"
)

var ErrRepository = errors.New("list activities error repository")

func TestListAccountActivities(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name string
		repo AcitivitiesMock
		args string
		want []AccountActivity
		err  error
	}

	wantActitivies := []AccountActivity{
		{
			Type:      Bill,
			Amount:    132000,
			CreatedAt: time.Now(),
			Details: DescriptionPayment{
				Description: "Academia",
				Status:      bills.Pago,
			},
		},
		{
			Type:      Bill,
			Amount:    560000,
			CreatedAt: time.Now(),
			Details: DescriptionPayment{
				Description: "IPTU",
				Status:      bills.Agendado,
			},
		},
		{
			Type:      Bill,
			Amount:    30000,
			CreatedAt: time.Now(),
			Details: DescriptionPayment{
				Description: "Escola",
				Status:      bills.Negado,
			},
		},
		{
			Type:      Transfer,
			Amount:    65222,
			CreatedAt: time.Now(),
			Details: DestinyAccount{
				Name:                 "Thomas",
				AccountDestinationId: uuid.NewString(),
			},
		},
	}

	id := uuid.NewString()

	testCases := []TestCase{
		{
			Name: "Should list activities successful",
			repo: AcitivitiesMock{
				OnListActivities: func(accountId string) ([]AccountActivity, error) {
					return wantActitivies, nil
				},
			},
			args: id,
			want: wantActitivies,
		},
		{
			Name: "Fail if return error from repository",
			repo: AcitivitiesMock{
				OnListActivities: func(accountId string) ([]AccountActivity, error) {
					return []AccountActivity{}, ErrRepository
				},
			},
			args: id,
			want: []AccountActivity{},
			err:  ErrRepository,
		},
	}

	for _, tc := range testCases {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			uc := NewAccountActivityUsecase(tt.repo)

			got, err := uc.ListActivity(context.Background(), tt.args)

			if !errors.Is(tt.err, err) {
				t.Errorf("Expected error %s, got %s", tt.err, err)
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Expected %v, got %v", tt.want, got)
			}
		})
	}
}
