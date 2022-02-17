package activities

import (
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/entities/bills"
)

type TypeActivity string

const (
	Transfer TypeActivity = "Transferência"
	Bill     TypeActivity = "Pagamento"
)

type AccountActivity struct {
	Type      TypeActivity
	Amount    int
	CreatedAt time.Time
	Details   interface{}
}

type DestinyAccount struct {
	AccountDestinationId string
	Name                 string
}

type DescriptionPayment struct {
	Description string
	Status      bills.Status
}
