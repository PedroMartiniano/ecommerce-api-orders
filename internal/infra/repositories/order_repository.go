package repositories

import (
	"context"
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) ports.IOrderRepository {
	return OrderRepository{
		db: db,
	}
}

func (r OrderRepository) SaveOrder(c context.Context, order entities.Order) (entities.Order, error) {
	query := `
		INSERT INTO orders (id, user_id, order_status, total_amount, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6)	
	`
	_, err := r.db.ExecContext(
		c,
		query,
		order.ID,
		order.UserID,
		order.OrderStatus,
		order.TotalAmount,
		order.CreatedAt,
		order.UpdatedAt,
	)
	if err != nil {
		return entities.Order{}, configs.NewError(configs.ErrBadRequest, err)
	}

	return order, nil
}
