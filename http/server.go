package http

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/usecases"
	"github.com/gorilla/mux"
)

type Error struct {
	Reason string `json:"reason"`
}

type Server struct {
	app usecases.Accounts
	tr  usecases.Tranfers
	http.Handler
}

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

func NewServer(
	usecaseAcc usecases.Accounts,
	usecaseTr usecases.Tranfers,
) Server {

	server := Server{
		app: usecaseAcc,
		tr:  usecaseTr,
	}

	router := mux.NewRouter()

	router.HandleFunc("/accounts", server.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts", server.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account_id}/balance", server.ListOne).Methods(http.MethodGet)
	router.HandleFunc("/login", server.Login).Methods((http.MethodPost))
	router.HandleFunc("/transfers", server.ListTransfer).Methods((http.MethodGet))
	router.HandleFunc("/transfers", server.CreateTransfer).Methods((http.MethodPost))

	server.Handler = router
	return server
}
