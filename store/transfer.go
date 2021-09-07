package store

import "time"

type Transfer struct {
	Id                   string    `json:"id,omitempty"`
	AccountOriginId      string    `json:"account_origin_id,omitempty"`
	AccountDestinationId string    `json:"account_destination_id,omitempty"`
	Amount               float64   `json:"amount,omitempty"`
	createdAt            time.Time `json:"created_at"`
}