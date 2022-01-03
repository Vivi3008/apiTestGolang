package bills

type BillRepository interface {
	ListBills(string) ([]Bill, error)
	StoreBill(Bill) error
}
