package dtos

type ProcessPaymentDTO struct {
	OrderID        string  `json:"order_id"`
	CardHolder     string  `json:"card_holder"`
	CardNumber     string  `json:"card_number"`
	ExpirationDate string  `json:"expiration_date"`
	CVV            string  `json:"cvv"`
	Amount         float64 `json:"amount"`
}
