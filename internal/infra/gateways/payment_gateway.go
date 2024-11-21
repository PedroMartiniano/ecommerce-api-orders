package gateways

import (
	pg "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/gateways"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/queue"
)

type PaymentGateway struct{
	queue queue.Queue
}

func NewPaymentGateway(queue queue.Queue) pg.IPaymentGateway {
	return &PaymentGateway{
		queue: queue,
	}
}

func (p *PaymentGateway) ProcessPayment(dto dtos.ProcessPaymentDTO) error {
	return p.queue.PublishMessage("process-payment", dto)
}
