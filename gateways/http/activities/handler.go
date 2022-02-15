package activities

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/activities"
	"github.com/gorilla/mux"
)

type Handler struct {
	actUse activities.ActivityUsecase
}

func NewHandler(router *mux.Router, actUse activities.ActivityUsecase) *Handler {
	h := &Handler{actUse: actUse}

	router.HandleFunc("/activity", h.ListActivity).Methods(http.MethodGet)
	return h
}
