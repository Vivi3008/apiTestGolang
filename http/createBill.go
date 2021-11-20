package http

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Vivi3008/apiTestGolang/domain"
	"github.com/dgrijalva/jwt-go"
)

type BillReqRes struct {
	AccountId     string        `json:"account_id"`
	Description   string        `json:"description"`
	Value         float64       `json:"value"`
	DueDate       time.Time     `json:"due_date"`
	ScheduledDate time.Time     `json:"scheduled_date"`
	StatusBill    domain.Status `json:"status"`
}

func (s Server) CreateBill(w http.ResponseWriter, r *http.Request) {
	if r.Header["Authorization"] == nil {
		response := Error{Reason: "Auth required"}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusUnauthorized)
		json.NewEncoder(w).Encode(response)
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

	var body BillReqRes

	err = json.NewDecoder(r.Body).Decode(&body)

	if err != nil {
		response := Error{Reason: "invalid request body"}
		log.Printf("error decoding body: %s\n", err.Error())
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	bill := domain.Bill{
		AccountId:   accountId,
		Description: body.Description,
		Value:       body.Value,
		DueDate:     body.DueDate,
	}

	billOk, err := s.app.CreateBill(bill)

	if err != nil {
		log.Printf("Failed to pay bill: %s\n", err.Error())
		response := Error{Reason: err.Error()}
		w.Header().Set(ContentType, JSONContentType)
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(response)
		return
	}

	billStatusOk := domain.Bill{
		Id:          billOk.Id,
		AccountId:   accountId,
		Description: billOk.Description,
		Value:       billOk.Value,
		DueDate:     billOk.DueDate,
		StatusBill:  domain.Pago,
	}

	saveBill, err := s.bl.SaveBill(billStatusOk)

	if err != nil {
		log.Printf("Failed to save bill: %s\n", err.Error())
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
		return
	}

	billResonse := BillReqRes{
		AccountId:     saveBill.AccountId,
		Description:   saveBill.Description,
		Value:         saveBill.Value,
		DueDate:       saveBill.DueDate,
		ScheduledDate: saveBill.ScheduledDate,
		StatusBill:    saveBill.StatusBill,
	}

	w.Header().Set(ContentType, JSONContentType)
	json.NewEncoder(w).Encode(billResonse)
	log.Printf("sent successful response for transfer %s\n", saveBill.Id)
}
