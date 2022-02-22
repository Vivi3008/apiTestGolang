package transfers

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/gorilla/mux"
)

type Handler struct {
	transfUse transfers.Usecase
	accUse    account.Usecase
}

func NewHandler(router *mux.Router, usecase transfers.Usecase, accUsecase account.Usecase) *Handler {
	h := &Handler{
		transfUse: usecase,
		accUse:    accUsecase,
	}

	router.HandleFunc("/transfers", h.CreateTransfer).Methods(http.MethodPost)
	router.HandleFunc("/transfers", h.ListTransfer).Methods(http.MethodGet)
	return h
}
