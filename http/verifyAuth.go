package http

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/dgrijalva/jwt-go"
)

var (
	ErrAuth = errors.New("authentication required")
)

func VerifyAuth(w http.ResponseWriter, r *http.Request) (domain.AccountId, error) {
	if r.Header["Authorization"] == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Authentication required")
		return "", ErrAuth
	}

	authHeader := r.Header.Get("Authorization")

	var accountId string

	//pegar o id do token
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountId = claims["id"].(string)
	} else {
		res := err.Error()
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(res)
		return "", ErrAuth
	}

	return domain.AccountId(accountId), nil
}
