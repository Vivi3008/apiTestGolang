package activities

import (
	"errors"
	"log"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
)

var ErrGetTokenId = errors.New("error to get id from token")

type ActivitiesResponse struct {
	Type      activities.TypeActivity `json:"type"`
	Amount    int                     `json:"amount"`
	CreatedAt time.Time               `json:"created_at"`
	Details   interface{}             `json:"details"`
}

func (h Handler) ListActivity(w http.ResponseWriter, r *http.Request) {
	accountId, ok := middlewares.GetAccountId(r.Context())

	if !ok || accountId == "" {
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	listActivities, err := h.actUse.ListActivity(r.Context(), accountId)
	if err != nil {
		log.Printf("Failed to list activities: %s\n", err.Error())
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	listResponse := make([]ActivitiesResponse, len(listActivities))

	for i, ls := range listActivities {
		listResponse[i].Type = ls.Type
		listResponse[i].Amount = ls.Amount
		listResponse[i].CreatedAt = ls.CreatedAt
		listResponse[i].Details = ls.Details
	}

	response.Send(w, listResponse, http.StatusOK)
	log.Printf("Sent all activities from Id %s", accountId)
	log.Printf("Sent all activities. Total: %d", len(listResponse))
}
