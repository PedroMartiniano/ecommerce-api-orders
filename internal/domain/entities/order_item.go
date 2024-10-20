package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/vo"
)

type OrderItem struct {
	ID         vo.UUID     `json:"id"`
	OrderID    vo.UUID     `json:"order_id"`
	ProductID  vo.UUID     `json:"product_id"`
	Quantity   vo.Quantity `json:"quantity"`
	UnitPrice  vo.Amount   `json:"price"`
	TotalPrice vo.Amount   `json:"total"`
	CreatedAt  time.Time   `json:"created_at"`
}
