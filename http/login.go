package http

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/dgrijalva/jwt-go"
)

type LoginRequest struct {
	Cpf    int64  `json:"cpf"`
	Secret string `json:"secret"`
}

type TokenString struct {
	Token string `json:"token"`
}

const ACCESS_SECRET = "fadsfasf6s5f65sa6"

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

	login := domain.Login{
		Cpf:    body.Cpf,
		Secret: body.Secret,
	}

	accountId, err := s.app.NewLogin(login)

	if accountId == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode("Cpf invalid")
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

func createToken(accountId domain.AccountId) (string, error) {
	idClaims := jwt.MapClaims{}
	idClaims["id"] = accountId

	tokenStr := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), idClaims)

	tokenString, err := tokenStr.SignedString([]byte(ACCESS_SECRET))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}
