package auth

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/gorilla/mux"
)

type Handler struct {
	accUse account.AccountUsecase
}

func NewHandler(router *mux.Router, usecase account.AccountUsecase) *Handler {
	h := &Handler{accUse: usecase}

	router.HandleFunc("/login", h.Login).Methods((http.MethodPost))
	return h
}
