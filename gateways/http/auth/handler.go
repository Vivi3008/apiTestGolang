package auth

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/gorilla/mux"
)

type Handler struct {
	accUse account.Usecase
}

func NewHandler(router *mux.Router, usecase account.Usecase) *Handler {
	h := &Handler{accUse: usecase}

	router.HandleFunc("/login", h.Login).Methods(http.MethodPost)
	return h
}
