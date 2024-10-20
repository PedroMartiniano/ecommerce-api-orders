package dto

type OrderDTO struct {
	ID          string
	UserID      string
	Quantity    int
	OrderStatus string
	TotalAmount float64
	CreatedAt   string
	UpdatedAt   string
}

type OrderItemDTO struct {
	ID         string
	OrderID    string
	ProductID  string
	Quantity   int
	UnitPrice  float64
	TotalPrice float64
	CreatedAt  string
}
