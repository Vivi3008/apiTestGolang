package http

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/Vivi3008/apiTestGolang/http/accounts"
	"github.com/Vivi3008/apiTestGolang/http/auth"
	"github.com/Vivi3008/apiTestGolang/http/middlewares"
	"github.com/gorilla/mux"
)

type Server struct {
	app account.AccountUsecase
	tr  transfers.TranfersUsecase
	bl  bill.BillUsecase
	http.Handler
}

func NewServer(
	accountUc account.AccountUsecase,
	usecaseTr transfers.TranfersUsecase,
	usecaseBl bill.BillUsecase,
) Server {

	server := Server{
		tr: usecaseTr,
		bl: usecaseBl,
	}

	router := mux.NewRouter()
	routerAuth := router.NewRoute().Subrouter()

	accounts.NewHandler(router, accountUc)
	auth.NewHandler(router, accountUc)

	routerAuth.HandleFunc("/transfers", server.CreateTransfer).Methods((http.MethodPost))
	routerAuth.HandleFunc("/bills", server.CreateBill).Methods((http.MethodPost))
	routerAuth.HandleFunc("/bills", server.ListBills).Methods((http.MethodGet))
	routerAuth.HandleFunc("/transfers", server.ListTransfer).Methods(http.MethodGet)
	routerAuth.Use(middlewares.Auth)

	server.Handler = router
	return server
}
