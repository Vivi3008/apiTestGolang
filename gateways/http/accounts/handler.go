package accounts

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/gorilla/mux"
)

type Handler struct {
	acc account.Usecase
}

func NewHandler(router *mux.Router, accUse account.Usecase) *Handler {
	h := &Handler{acc: accUse}

	router.HandleFunc("/accounts", h.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts", h.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account_id}/balance", h.GetBalance).Methods(http.MethodGet)
	return h
}
