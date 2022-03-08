package activities

import (
	"errors"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/Vivi3008/apiTestGolang/gateways/http/middlewares"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	lg "github.com/Vivi3008/apiTestGolang/infraestructure/logging"
)

var ErrGetTokenId = errors.New("error to get id from token")

type ActivitiesResponse struct {
	Type      activities.TypeActivity `json:"type"`
	Amount    int                     `json:"amount"`
	CreatedAt time.Time               `json:"created_at"`
	Details   interface{}             `json:"details"`
}

func (h Handler) ListActivity(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.activity.ListActivity"
	accountId, ok := middlewares.GetAccountId(r.Context())

	log := lg.FromContext(r.Context(), operation)

	if !ok || accountId == "" {
		log.Error("Error to get token id for account")
		response.SendError(w, ErrGetTokenId, http.StatusUnauthorized)
		return
	}

	listActivities, err := h.actUse.ListActivity(r.Context(), accountId)
	if err != nil {
		log.Error("Failed to list activities: ", err)
		response.SendError(w, err, http.StatusInternalServerError)
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
	log.WithField("accountId", accountId).Info("Sent all activities from account")
	log.WithField("Total", len(listResponse)).Info("List total activities")
}
