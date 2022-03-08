package accounts

import (
	"errors"
	"net/http"

	lg "github.com/Vivi3008/apiTestGolang/gateways/http/logging"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
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

	log := lg.FromContext(r.Context(), operation)

	log.Info("Starting to get all accounts")
	list, err := h.acc.ListAllAccounts(r.Context())

	if err != nil {
		log.WithError(err).Error("Error to list all accounts")
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

	log.Info("Sent all accounts. Total: ", len(accounts))
	response.Send(w, accounts, http.StatusOK)
}

func (h Handler) GetBalance(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.account.GetBalance"
	vars := mux.Vars(r)

	accountId := vars["account_id"]

	log := lg.FromContext(r.Context(), operation)
	log.WithField("accountId", accountId).Info("Starting to get balance")

	account, err := h.acc.ListAccountById(r.Context(), accountId)

	if err != nil {
		log.WithError(err).Error("Failed to list account")
		response.SendError(w, err, http.StatusInternalServerError)
		return
	}

	balance := BalanceAccountResponse{
		Balance: account.Balance,
	}

	log.WithField("accountId", accountId).Info("Get balance sucessfull")
	response.Send(w, balance, http.StatusOK)
}
