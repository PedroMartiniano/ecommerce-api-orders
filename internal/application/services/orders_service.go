package services

import (
	"context"
	"errors"

	pg "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/gateways"
	pr "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/repositories"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/entities"
)

var logger = configs.GetLogger()

type OrdersService struct {
	ordersRepository     pr.IOrderRepository
	orderItemsRepository pr.IOrderItemsRepository
	paymentGateway       pg.IPaymentGateway
	productGateway       pg.IProductGateway
}

func NewOrdersService(ordersRepository pr.IOrderRepository, orderItemsRepository pr.IOrderItemsRepository, paymentGateway pg.IPaymentGateway, productGateway pg.IProductGateway) *OrdersService {
	return &OrdersService{
		ordersRepository:     ordersRepository,
		orderItemsRepository: orderItemsRepository,
		paymentGateway:       paymentGateway,
		productGateway:       productGateway,
	}
}

func (s *OrdersService) CreateOrderExecute(c context.Context, orderDTO dtos.OrderDTO) (dtos.OrderResDTO, error) {
	order, err := entities.CreateNewOrder(orderDTO.UserID, orderDTO.AddressID, orderDTO.TotalAmount)
	if err != nil {
		return dtos.OrderResDTO{}, configs.NewError(configs.ErrBadRequest, err)
	}

	orderItems := []entities.OrderItem{}
	for _, item := range orderDTO.Items {
		product, err := s.productGateway.GetProductByID(item.ProductID)
		if err != nil {
			order.FailOrder()
			s.ordersRepository.SaveOrder(c, order)
			return dtos.OrderResDTO{}, configs.NewError(configs.ErrBadRequest, err)
		}

		if product.Price != item.UnitPrice {
			item.UnitPrice = product.Price
		}

		if product.Quantity < item.Quantity {
			order.FailOrder()
			s.ordersRepository.SaveOrder(c, order)
			return dtos.OrderResDTO{}, configs.NewError(configs.ErrBadRequest, errors.New("quantidade do produto '"+product.Name+"' insuficiente"))
		}

		item.TotalPrice = float64(item.Quantity) * item.UnitPrice
		orderItem, err := entities.CreateNewOrderItem(order.GetID(), item.ProductID, item.Quantity, item.UnitPrice, item.TotalPrice)
		if err != nil {
			return dtos.OrderResDTO{}, configs.NewError(configs.ErrBadRequest, err)
		}
		orderItems = append(orderItems, orderItem)
	}

	err = s.paymentGateway.ProcessPayment(dtos.ProcessPaymentDTO{
		OrderID:        order.GetID(),
		Amount:         order.GetTotalAmount(),
		UserID:         orderDTO.UserID,
		CardHolder:     orderDTO.PaymentDetails.CardHolder,
		CardNumber:     orderDTO.PaymentDetails.CardNumber,
		ExpirationDate: orderDTO.PaymentDetails.ExpirationDate,
		CVV:            orderDTO.PaymentDetails.CVV,
	})
	if err != nil {
		order.FailOrder()
		s.ordersRepository.SaveOrder(c, order)
		return dtos.OrderResDTO{}, err
	}

	for _, item := range orderDTO.Items {
		err = s.productGateway.UpdateProductQuantity(dtos.UpdateProductQuantityDTO{
			Token:     orderDTO.Token,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Operation: "remove",
		})
		if err != nil {
			order.FailOrder()
			s.ordersRepository.SaveOrder(c, order)
			return dtos.OrderResDTO{}, err
		}
	}

	order, err = s.ordersRepository.SaveOrder(c, order)
	if err != nil {
		return dtos.OrderResDTO{}, err
	}

	_, err = s.orderItemsRepository.SaveOrderItems(c, orderItems)
	if err != nil {
		return dtos.OrderResDTO{}, err
	}

	return dtos.OrderResDTO{
		ID:          order.GetID(),
		UserID:      order.GetUserID(),
		AddressID:   order.GetAddressID(),
		TotalAmount: order.GetTotalAmount(),
		Status:      order.GetOrderStatus(),
		CreatedAt:   order.GetCreatedAt(),
		UpdatedAt:   order.GetUpdatedAt(),
	}, nil
}
