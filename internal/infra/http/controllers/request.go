package controllers

type CreateOrderRequest struct {
	AddressID   string  `json:"address_id"`
	TotalAmount float64 `json:"total_amount"`
	Items       []struct {
		ProductID  string  `json:"product_id"`
		Quantity   int     `json:"quantity"`
		UnitPrice  float64 `json:"unit_price"`
		TotalPrice float64 `json:"total_price"`
	} `json:"items"`
	PaymentDetails struct {
		CardHolder     string `json:"card_holder"`
		CardNumber     string `json:"card_number"`
		ExpirationDate string `json:"expiration_date"`
		CVV            string `json:"cvv"`
	} `json:"payment_details"`
}
