package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/dgrijalva/jwt-go"
)

type TransferRequest struct {
	AccountDestinationId string  `json:"account_Destination_Id"`
	Amount               float64 `json:"amount"`
}

func (s Server) CreateTransfer(w http.ResponseWriter, r *http.Request) {
	if r.Header["Auth"] == nil {
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode("Authentication required")
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
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	saveTransfer, err := s.tr.SaveTransfer(transfer)

	if err != nil {
		log.Printf("Failed to save transfer: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	response := TransferResponse{
		Id:                   saveTransfer.Id,
		AccountOriginId:      saveTransfer.AccountOriginId,
		AccountDestinationId: saveTransfer.AccountDestinationId,
		Amount:               saveTransfer.Amount,
		CreatedAt:            saveTransfer.CreatedAt,
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("sent successful response for transfer %s\n", transfer.Id)
}
