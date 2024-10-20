package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/vo"
)

type Order struct {
	ID          vo.UUID   `json:"id"`
	UserID      vo.UUID   `json:"user_id"`
	OrderStatus string    `json:"order_status"`
	TotalAmount vo.Amount `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
