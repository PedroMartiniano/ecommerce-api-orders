package ports

import "github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"

type IPaymentGateway interface {
	ProcessPayment(dtos.ProcessPaymentDTO) error
}
