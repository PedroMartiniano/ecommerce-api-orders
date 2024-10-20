package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/vo"
)

type OrderItem struct {
	ID        vo.UUID   `json:"id"`
	OrderID   vo.UUID   `json:"order_id"`
	ProductID vo.UUID   `json:"product_id"`
	Quantity  int       `json:"quantity"`
	Price     vo.Amount `json:"price"`
	Total     vo.Amount `json:"total"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
