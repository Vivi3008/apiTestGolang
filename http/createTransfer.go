package http

type TransferRequest struct {
	AccountOriginId      string  `json:"account_origin_id"`
	AccountDestinationId string  `json:"account_Destination_Id"`
	Amount               float64 `json:"amount"`
}
