package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/dgrijalva/jwt-go"
)

type TransferResponse struct {
	Id                   string    `json:"id"`
	AccountOriginId      string    `json:"accoundId"`
	AccountDestinationId string    `json:"destinyId"`
	Amount               float64   `json:"amount"`
	CreatedAt            time.Time `json:"createdAt"`
}

func (s Server) ListTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Header["Authorization"] == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Authentication required")
		return
	}

	authHeader := r.Header.Get("Authorization")

	var accountId string

	//pegar o id do token
	token, err := jwt.Parse(authHeader, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method")
		}
		return []byte(ACCESS_SECRET), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		accountId = claims["id"].(string)
	} else {
		res := err.Error()
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(res)
		return
	}

	list, err := s.tr.ListTransfer(domain.AccountId(accountId))

	if err != nil {
		log.Printf("Failed to do list transfers: %s\n", err.Error())
		return
	}

	response := make([]TransferResponse, len(list))

	for i, transfer := range list {
		response[i].Id = transfer.Id
		response[i].AccountOriginId = transfer.AccountOriginId
		response[i].AccountDestinationId = transfer.AccountDestinationId
		response[i].Amount = transfer.Amount
		response[i].CreatedAt = transfer.CreatedAt
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("Sent all transfers from Id %s", accountId)
	log.Printf("Sent all transfers. Total: %d", len(response))
}
