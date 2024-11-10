package repositories

import (
	"context"
	"database/sql"

	pr "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

type OrderRepository struct {
	db *sql.DB
}

func NewOrderRepository(db *sql.DB) pr.IOrderRepository {
	return OrderRepository{
		db: db,
	}
}

func (r OrderRepository) SaveOrder(c context.Context, order entities.Order) (entities.Order, error) {
	query := `
		INSERT INTO orders (id, user_id, order_status, total_amount, created_at, updated_at, address_id)
		VALUES ($1, $2, $3, $4, $5, $6, $7)	
	`
	_, err := r.db.ExecContext(
		c,
		query,
		order.GetID(),
		order.GetUserID(),
		order.GetOrderStatus(),
		order.GetTotalAmount(),
		order.GetCreatedAt(),
		order.GetUpdatedAt(),
		order.GetAddressID(),
	)
	if err != nil {
		return entities.Order{}, configs.NewError(configs.ErrBadRequest, err)
	}

	return order, nil
}
