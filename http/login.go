package http

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain/entities/account"
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	Cpf    string `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenString struct {
	Token string `json:"token"`
}

func (s Server) Login(w http.ResponseWriter, r *http.Request) {
	var body LoginRequest

	err := json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	login := account.Login{
		Cpf:    body.Cpf,
		Secret: body.Secret,
	}

	accountId, err := s.app.NewLogin(login)

	if accountId == "" {
		response := Error{Reason: "Cpf does not exist"}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	if err != nil {
		response := Error{Reason: err.Error()}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	tokenString, err := createToken(accountId)

	if err != nil {
		response := Error{Reason: "Invalid token"}
		log.Printf("error invalid token: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	resToken := TokenString{
		Token: tokenString,
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(resToken)
}

func createToken(accountId string) (string, error) {
	idClaims := jwt.MapClaims{}
	idClaims["id"] = accountId

	tokenStr := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), idClaims)

	tokenString, err := tokenStr.SignedString([]byte(os.Getenv("ACCESS_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
