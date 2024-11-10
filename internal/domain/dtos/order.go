package dtos

type OrderDTO struct {
	Token          string         `json:"token"`
	UserID         string         `json:"user_id"`
	AddressID      string         `json:"address_id"`
	TotalAmount    float64        `json:"total_amount"`
	Items          []OrderItemDTO `json:"items"`
	PaymentDetails PaymentDetails `json:"payment_details"`
	CreatedAt      string         `json:"created_at"`
	UpdatedAt      string         `json:"updated_at"`
}

type OrderItemDTO struct {
	ID         string  `json:"id"`
	OrderID    string  `json:"order_id"`
	ProductID  string  `json:"product_id"`
	Quantity   int     `json:"quantity"`
	UnitPrice  float64 `json:"unit_price"`
	TotalPrice float64 `json:"total_price"`
	CreatedAt  string  `json:"created_at"`
}

type PaymentDetails struct {
	CardHolder     string `json:"card_holder"`
	CardNumber     string `json:"card_number"`
	ExpirationDate string `json:"expiration_date"`
	CVV            string `json:"cvv"`
}
