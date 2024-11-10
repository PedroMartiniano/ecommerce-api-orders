package gateways

import (
	pg "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/gateways"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"
)

type PaymentGateway struct{}

func NewPaymentGateway() pg.IPaymentGateway {
	return &PaymentGateway{}
}

func (p *PaymentGateway) ProcessPayment(dto dtos.ProcessPaymentDTO) error {
	return nil
}
