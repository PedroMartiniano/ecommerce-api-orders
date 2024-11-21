package adapters

import (
	"database/sql"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/gateways"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/queue"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/repositories"
)

func NewOrderServiceAdapter(db *sql.DB, queue queue.Queue) *services.OrdersService {
	ordersRepository := repositories.NewOrderRepository(db)
	orderItemsRepository := repositories.NewOrderItemsRepository(db)
	productGateway := gateways.NewProductGateway()
	paymentGateway := gateways.NewPaymentGateway(queue)

	return services.NewOrdersService(ordersRepository, orderItemsRepository, paymentGateway, productGateway)
}
