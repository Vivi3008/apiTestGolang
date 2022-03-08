package accounts

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	lg "github.com/Vivi3008/apiTestGolang/infraestructure/logging"
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

var ErrInvalidPayloadAccount = errors.New("invalid account payload")

func (h Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.account.CreateAccount"
	var body AccountRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.SendError(w, ErrInvalidPayloadAccount, http.StatusBadRequest)
		return
	}

	person := account.Account{
		Name:    body.Name,
		Cpf:     body.Cpf,
		Secret:  body.Secret,
		Balance: body.Balance,
	}

	log := lg.NewLog(r.Context(), operation)
	log.Info("Starting to create an account")

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
	log.Info("Create account successful with id: ", account.Id)
}
