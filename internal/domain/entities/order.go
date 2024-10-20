package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/vo"
)

type Order struct {
	ID          vo.UUID        `json:"id"`
	UserID      vo.UUID        `json:"user_id"`
	OrderStatus vo.OrderStatus `json:"order_status"`
	TotalAmount vo.Amount      `json:"total_amount"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}

func CreateNewOrder(userID string, orderStatus string, totalAmount float64) (Order, error) {
	now := time.Now()

	amount, err := vo.NewAmount(totalAmount)
	if err != nil {
		return Order{}, err
	}

	return Order{
		ID:          vo.NewUUID(""),
		UserID:      vo.NewUUID(userID),
		OrderStatus: vo.NewOrderStatus(vo.OrderStatusPending),
		TotalAmount: amount,
		CreatedAt:   now,
		UpdatedAt:   now,
	}, nil
}

func InstantiateOrder(id string, userID string, orderStatus string, totalAmount float64, createdAt time.Time, updatedAt time.Time) (Order, error) {
	amount, err := vo.NewAmount(totalAmount)
	if err != nil {
		return Order{}, err
	}

	return Order{
		ID:          vo.NewUUID(id),
		UserID:      vo.NewUUID(userID),
		OrderStatus: vo.NewOrderStatus(vo.Status(orderStatus)),
		TotalAmount: amount,
		CreatedAt:   createdAt,
		UpdatedAt:   updatedAt,
	}, nil
}
