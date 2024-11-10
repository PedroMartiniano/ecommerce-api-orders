package repositories

import (
	"context"
	"database/sql"

	pr "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

type OrderItemsRepository struct {
	db *sql.DB
}

func NewOrderItemsRepository(db *sql.DB) pr.IOrderItemsRepository {
	return OrderItemsRepository{
		db: db,
	}
}

func (r OrderItemsRepository) SaveOrderItems(c context.Context, orderItems []entities.OrderItem) ([]entities.OrderItem, error) {
	query := `
		INSERT INTO order_items (id, order_id, product_id, quantity, unit_price, total_price, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
	`

	stmt, err := r.db.PrepareContext(c, query)
	if err != nil {
		return []entities.OrderItem{}, configs.NewError(configs.ErrBadRequest, err)
	}
	defer stmt.Close()

	for _, item := range orderItems {
		_, err := stmt.ExecContext(
			c,
			item.GetID(),
			item.GetOrderID(),
			item.GetProductID(),
			item.GetQuantity(),
			item.GetUnitPrice(),
			item.GetTotalPrice(),
			item.GetCreatedAt(),
		)
		if err != nil {
			return []entities.OrderItem{}, configs.NewError(configs.ErrBadRequest, err)
		}
	}

	return orderItems, nil
}
