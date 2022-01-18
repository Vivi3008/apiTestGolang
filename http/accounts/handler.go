package accounts

import (
	"net/http"

	usecase "github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/gorilla/mux"
)

type Handler struct {
	acc usecase.AccountUsecase
}

func NewHandler(router *mux.Router, accUse usecase.AccountUsecase) *Handler {
	h := &Handler{accUse}

	router.HandleFunc("/accounts", h.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts", h.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/login", h.Login).Methods((http.MethodPost))
	router.HandleFunc("/accounts/{account_id}/balance", h.ListOne).Methods(http.MethodGet)
	return h
}
