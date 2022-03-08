package auth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/commom"
	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/Vivi3008/apiTestGolang/gateways/http/response"
	lg "github.com/Vivi3008/apiTestGolang/infraestructure/logging"
)

var ErrCpfNotExists = errors.New("cpf doesn't exists")

type LoginRequest struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenString struct {
	Token string `json:"token"`
}

var ErrInvalidLoginPayload = errors.New("invalid login payload")

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	const operation = "handler.auth.Login"
	var body LoginRequest

	log := lg.NewLog(r.Context(), operation)
	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		log.Error("Error to login: ", err)
		response.SendError(w, ErrInvalidLoginPayload, http.StatusBadRequest)
		return
	}

	login := account.Login{
		Cpf:    body.Cpf,
		Secret: body.Secret,
	}

	accountId, err := h.accUse.NewLogin(r.Context(), login)

	if err != nil {
		log.Error("Error to login: ", err)
		response.SendError(w, err, http.StatusUnauthorized)
		return
	}

	tokenString, err := commom.CreateToken(accountId)

	if err != nil {
		log.Error("Error to create token: ", err)
		response.SendError(w, err, http.StatusUnauthorized)
		return
	}

	resToken := TokenString{
		Token: tokenString,
	}
	log.Info("Login sucessfull for account: ", accountId)
	response.Send(w, resToken, http.StatusOK)
}
