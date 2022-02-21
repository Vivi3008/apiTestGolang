package bills

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/gorilla/mux"
)

type Handler struct {
	accUse account.Usecase
	blUse  bill.Usecase
}

func NewHandler(router *mux.Router, blUsecase bill.Usecase, acUsecase account.Usecase) *Handler {
	h := &Handler{
		accUse: acUsecase,
		blUse:  blUsecase,
	}

	router.HandleFunc("/bills", h.CreateBill).Methods(http.MethodPost)
	router.HandleFunc("/bills", h.ListBills).Methods(http.MethodGet)

	return h
}
