package transfers

type TransferRepository interface {
	ListTransfer(string) ([]Transfer, error)
	SaveTransfer(Transfer) error
	CreateTransfer(Transfer) (Transfer, error)
}
