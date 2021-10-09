package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/dgrijalva/jwt-go"
)

type TransferRequest struct {
	AccountDestinationId string  `json:"account_Destination_Id"`
	Amount               float64 `json:"amount"`
}

func (s Server) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	validAuth := strings.Split(r.Header.Get("Auth"), ",")

	if len(validAuth) == 0 {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Invalid Token")
		return
	}

	authHeader := r.Header.Get("Auth")

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
		return
	}

	var body TransferRequest

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	transaction := domain.Transfer{
		AccountOriginId:      accountId,
		AccountDestinationId: body.AccountDestinationId,
		Amount:               body.Amount,
	}

	transfer, err := s.app.CreateTransfer(transaction)

	if err != nil {
		log.Printf("Failed to do transfer: %s\n", err.Error())
		return
	}

	response := TransferResponse{
		Id:                   transfer.Id,
		AccountOriginId:      transfer.AccountOriginId,
		AccountDestinationId: transfer.AccountDestinationId,
		Amount:               transfer.Amount,
		CreatedAt:            transfer.CreatedAt,
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent successful response for transfer %s\n", transfer.Id)
}
