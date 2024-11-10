package ports

import (
	"context"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

type IOrderRepository interface {
	SaveOrder(context.Context, entities.Order) (entities.Order, error)
}
