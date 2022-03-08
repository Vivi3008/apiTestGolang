package response

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Error struct {
	Reason string `json:"reason"`
}

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = time.RFC3339
)

func SendError(w http.ResponseWriter, err error, statusCode int) {
	response := Error{Reason: err.Error()}
	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(statusCode)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Fatalf("Error to encode response %s", err)
	}
}

func Send(w http.ResponseWriter, res interface{}, statusCode int) {
	w.Header().Set(ContentType, JSONContentType)
	w.WriteHeader(statusCode)
	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Fatalf("Error to encode response %s", err)
	}
}
