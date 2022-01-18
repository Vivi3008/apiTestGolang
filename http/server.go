package http

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases/account"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/bill"
	"github.com/Vivi3008/apiTestGolang/domain/usecases/transfers"
	"github.com/gorilla/mux"
)

type Error struct {
	Reason string `json:"reason"`
}

type Server struct {
	app account.AccountUsecase
	tr  transfers.TranfersUsecase
	bl  bill.BillUsecase
	http.Handler
}

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

func NewServer(
	usecaseAcc account.AccountUsecase,
	usecaseTr transfers.TranfersUsecase,
	usecaseBl bill.BillUsecase,
) Server {

	server := Server{
		app: usecaseAcc,
		tr:  usecaseTr,
		bl:  usecaseBl,
	}

	router := mux.NewRouter()
	routerAuth := router.NewRoute().Subrouter()

	router.HandleFunc("/accounts", server.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts", server.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/login", server.Login).Methods((http.MethodPost))

	routerAuth.HandleFunc("/accounts/{account_id}/balance", server.ListOne).Methods(http.MethodGet)
	routerAuth.HandleFunc("/transfers", server.CreateTransfer).Methods((http.MethodPost))
	routerAuth.HandleFunc("/bills", server.CreateBill).Methods((http.MethodPost))
	routerAuth.HandleFunc("/bills", server.ListBills).Methods((http.MethodGet))
	routerAuth.HandleFunc("/transfers", server.ListTransfer).Methods(http.MethodGet)
	routerAuth.Use(Auth)

	server.Handler = router
	return server
}
