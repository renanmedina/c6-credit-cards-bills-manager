package management

type PurchaseItem struct {
	Date         string
	ClientName   string
	CreditCardId int
	Category     string
	Description  string
	Installment  string
	Amount       float64
}

func NewPurchaseItem(date string, client string, cardId int, category string, description string, installment string, amount float64) PurchaseItem {
	return PurchaseItem{
		Date:         date,
		ClientName:   client,
		CreditCardId: cardId,
		Category:     category,
		Description:  description,
		Installment:  installment,
		Amount:       amount,
	}
}
