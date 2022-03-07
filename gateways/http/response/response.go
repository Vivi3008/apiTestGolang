package response

import (
	"encoding/json"
	"log"
	"net/http"
)

type Error struct {
	Reason string `json:"reason"`
}

const (
	ContentType     = "Content-Type"
	JSONContentType = "application/json"
	DateLayout      = "2006-01-02T15:04:05Z"
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
