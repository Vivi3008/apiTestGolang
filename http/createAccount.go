package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain"
)

type AccountRequest struct {
	Name    string `json:"name"`
	Cpf     int    `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int64  `json:"balance"`
}

type AccountResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Cpf       int    `json:"cpf"`
	Balance   int64  `json:"balance"`
	CreatedAt string `json:"createdAt"`
}

func (s Server) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var body AccountRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	person := domain.Account{
		Name:    body.Name,
		Cpf:     body.Cpf,
		Secret:  body.Secret,
		Balance: body.Balance,
	}

	account, err := s.app.CreateAccount(person)

	if err != nil {
		response := Error{Reason: err.Error()}
		log.Printf("Failed to create account: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := AccountResponse{
		Id:        account.Id,
		Name:      account.Name,
		Cpf:       account.Cpf,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt.Format(DateLayout),
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent successful response for account %s\n", account.Id)
}
