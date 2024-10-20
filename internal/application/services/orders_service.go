package services

import (
	"context"

	p "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dto"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

type OrdersService struct {
	ordersRepository     p.IOrderRepository
	orderItemsRepository p.IOrderItemsRepository
}

func NewOrdersService(ordersRepository p.IOrderRepository, orderItemsRepository p.IOrderItemsRepository) *OrdersService {
	return &OrdersService{
		ordersRepository:     ordersRepository,
		orderItemsRepository: orderItemsRepository,
	}
}

func (s *OrdersService) CreateOrder(c context.Context, orderDTO dto.OrderDTO) (dto.OrderDTO, error) {
	_, err := entities.CreateNewOrder(orderDTO.UserID, orderDTO.OrderStatus, orderDTO.TotalAmount)
	if err != nil {
		return dto.OrderDTO{}, configs.NewError(configs.ErrBadRequest, err)
	}

	return orderDTO, nil

}
