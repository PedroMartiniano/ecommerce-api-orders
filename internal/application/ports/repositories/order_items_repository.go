package ports

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

type IOrderItemsRepository interface {
	SaveOrderItems(context.Context, []entities.OrderItem) ([]entities.OrderItem, error)
}
