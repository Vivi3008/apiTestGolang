package accounts

import (
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/http/response"
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

func (h Handler) ListAll(w http.ResponseWriter, r *http.Request) {
	list, err := h.acc.ListAllAccounts()

	if err != nil {
		log.Printf("Failed to list accounts: %s\n", err.Error())
		response.SendError(w, err, http.StatusBadRequest)
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

	response.SendRequest(w, accounts, http.StatusOK)
	log.Printf("Sent all accounts. Total: %d", len(accounts))
}

func (h Handler) ListOne(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	personId := vars["account_id"]

	account, err := h.acc.ListAccountById(personId)

	if err != nil {
		log.Printf("Failed to list account: %s", err.Error())
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	balance := BalanceAccountResponse{
		Balance: account.Balance,
	}

	response.SendRequest(w, balance, http.StatusOK)
}
