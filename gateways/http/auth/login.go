package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
)

var ErrCpfNotExists = errors.New("cpf doesn't exists")

type LoginRequest struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenString struct {
	Token string `json:"token"`
}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response.SendError(w, err, http.StatusBadRequest)
		return
	}

	login := account.Login{
		Cpf:    body.Cpf,
		Secret: body.Secret,
	}

	accountId, err := h.accUse.NewLogin(login)

	if err != nil {
		response.SendError(w, err, http.StatusUnauthorized)
		return
	}

	tokenString, err := commom.CreateToken(accountId)

	if err != nil {
		response.SendError(w, err, http.StatusUnauthorized)
		return
	}

	resToken := TokenString{
		Token: tokenString,
	}

	response.Send(w, resToken, http.StatusOK)
}
