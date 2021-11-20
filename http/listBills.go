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

func (s Server) ListBills(w http.ResponseWriter, r *http.Request) {
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

	list, err := s.bl.ListAllBills(domain.AccountId(accountId))

	if err != nil {
		log.Printf("Failed to do list bills: %s\n", err.Error())
		return
	}

	response := make([]BillReqRes, len(list))

	for i, bill := range list {
		response[i].AccountId = bill.AccountId
		response[i].Description = bill.Description
		response[i].Value = bill.Value
		response[i].DueDate = bill.DueDate
		response[i].ScheduledDate = bill.ScheduledDate
		response[i].StatusBill = bill.StatusBill
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(response)
	log.Printf("Sent all bills from Id %s", accountId)
	log.Printf("Sent all bills. Total: %d", len(response))
}
