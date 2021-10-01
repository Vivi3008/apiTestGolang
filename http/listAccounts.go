package http

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain"
)

type ListAccountResponse struct {
	Id        string    `json:"id"`
	Name      string    `json:"name"`
	Cpf       int64     `json:"cpf"`
	Balance   float64   `json:"balance"`
	CreatedAt time.Time `json:"createdAt"`
}

type AccountIdRequest struct {
	Id string `json:"id"`
}

func (s Server) ListAll(w http.ResponseWriter, r *http.Request) {
	list, err := s.accounts.ListAllAccounts()

	if err != nil {
		log.Printf("Failed to list accounts: %s\n", err.Error())
		response := Error{Reason: "internal server error"}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(response)
		return
	}

	response := make([]ListAccountResponse, len(list))

	for i, account := range list {
		response[i].Id = account.Id
		response[i].Name = account.Name
		response[i].Cpf = account.Cpf
		response[i].Balance = account.Balance
		response[i].CreatedAt = account.CreatedAt
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("Sent all accounts. Total: %d", len(response))
}

func (s Server) ListOne(w http.ResponseWriter, r *http.Request) {
	var body AccountIdRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	personId := domain.AccountId{
		Id: body.Id,
	}

	account, err := s.accounts.ListAccountById(personId)

	response := ListAccountResponse{
		Id:        account.Id,
		Name:      account.Name,
		Cpf:       account.Cpf,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt,
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("Sent accounts with Id: %s", response.Id)
}
