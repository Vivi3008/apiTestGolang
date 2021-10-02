package http

import (
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/gorilla/mux"
)

type Error struct {
	Reason string `json:"reason"`
}

type Server struct {
	accounts domain.Usecase
	http.Handler
}

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
)

func NewServer(usecase domain.Usecase) Server {
	server := Server{accounts: usecase}

	router := mux.NewRouter()

	router.HandleFunc("/accounts", server.CreateAccount).Methods(http.MethodPost)
	router.HandleFunc("/accounts", server.ListAll).Methods(http.MethodGet)
	router.HandleFunc("/accounts/{account_id}/balance", server.ListOne).Methods(http.MethodGet)

	server.Handler = router
	return server
}
