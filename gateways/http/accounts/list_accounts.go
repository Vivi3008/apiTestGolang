package accounts

import (
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	lg "github.com/Vivi3008/apiTestGolang/infraestructure/logging"
	"github.com/gorilla/mux"
)

type ListAccountResponse struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Cpf       string `json:"cpf"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"createdAt"`
}

type BalanceAccountResponse struct {
	Balance int `json:"balance"`
}

type AccountIdRequest struct {
	Id string `json:"id"`
}

var ErrInvalidParam = errors.New("invalid id params")

func (h Handler) ListAll(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.account.ListAll"

	log := lg.NewLog(r.Context(), operation)
	log.Info("Starting to get all accounts")

	list, err := h.acc.ListAllAccounts(r.Context())

	if err != nil {
		log.Error("Error to list all accounts", err)
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	accounts := make([]ListAccountResponse, len(list))

	for i, account := range list {
		accounts[i].Id = account.Id
		accounts[i].Name = account.Name
		accounts[i].Cpf = account.Cpf
		accounts[i].Balance = account.Balance
		accounts[i].CreatedAt = account.CreatedAt.Format(response.DateLayout)
	}

	response.Send(w, accounts, http.StatusOK)
	log.Info("Sent all accounts. Total: ", len(accounts))
}

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.account.GetBalance"
	vars := mux.Vars(r)

	accountId := vars["account_id"]

	log := lg.NewLog(r.Context(), operation)
	log.Info("Starting to get balance for account id: ", accountId)

	account, err := h.acc.ListAccountById(r.Context(), accountId)

	if err != nil {
		log.Error("Failed to list account: ", err)
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	balance := BalanceAccountResponse{
		Balance: account.Balance,
	}

	log.Info("Get balance sucessfull for account id: ", accountId)
	response.Send(w, balance, http.StatusOK)
}
