package activities

import (
	"time"
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

type ActivityUsecase struct {
	actRepo AccountActivityRepository
}

func NewAccountActivityUsecase(act AccountActivityRepository) ActivityUsecase {
	return ActivityUsecase{actRepo: act}
}
