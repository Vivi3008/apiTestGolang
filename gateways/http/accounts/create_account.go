package accounts

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
)

type AccountRequest struct {
	Name    string `json:"name"`
	Cpf     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

type AccountResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"createdAt"`
}

func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var body AccountRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	person := account.Account{
		Name:    body.Name,
		Cpf:     body.Cpf,
		Secret:  body.Secret,
		Balance: body.Balance,
	}

	account, err := h.acc.CreateAccount(r.Context(), person)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	accountResponse := AccountResponse{
		Id:        account.Id,
		Name:      account.Name,
		Cpf:       account.Cpf,
		Balance:   account.Balance,
		CreatedAt: account.CreatedAt.Format(response.DateLayout),
	}

	response.Send(w, accountResponse, http.StatusOK)
	log.Printf("sent successful response for account %s\n", account.Id)
}
