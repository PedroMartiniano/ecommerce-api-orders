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

func (s *OrdersService) CreateOrderExecute(c context.Context, orderDTO dtos.OrderDTO) error {
	order, err := entities.CreateNewOrder(orderDTO.UserID, orderDTO.AddressID, orderDTO.TotalAmount)
	if err != nil {
		return configs.NewError(configs.ErrBadRequest, err)
	}

	orderItems := []entities.OrderItem{}
	for _, item := range orderDTO.Items {
		orderItem, err := entities.CreateNewOrderItem(order.GetID(), item.ProductID, item.Quantity, item.UnitPrice, item.TotalPrice)
		if err != nil {
			return err
		}
		orderItems = append(orderItems, orderItem)
	}

	for _, item := range orderDTO.Items {
		product, err := s.productGateway.GetProductByID(item.ProductID)
		if err != nil {
			order.FailOrder()
			s.ordersRepository.SaveOrder(c, order)
			return err
		}
		logger.Debugf("%+v", product)
		if product.Quantity < item.Quantity {
			order.FailOrder()
			s.ordersRepository.SaveOrder(c, order)
			return configs.NewError(configs.ErrBadRequest, errors.New("quantidade do produto '"+product.Name+"' insuficiente"))
		}

	}

	err = s.paymentGateway.ProcessPayment(dtos.ProcessPaymentDTO{
		OrderID:        order.GetID(),
		Amount:         order.GetTotalAmount(),
		CardHolder:     orderDTO.PaymentDetails.CardHolder,
		CardNumber:     orderDTO.PaymentDetails.CardNumber,
		ExpirationDate: orderDTO.PaymentDetails.ExpirationDate,
		CVV:            orderDTO.PaymentDetails.CVV,
	})
	if err != nil {
		order.FailOrder()
		s.ordersRepository.SaveOrder(c, order)
		return err
	}

	for _, item := range orderDTO.Items {	
		err = s.productGateway.UpdateProductQuantity(dtos.UpdateProductQuantityDTO{
			Token:       orderDTO.Token,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			Operation: "remove",
		})
		if err != nil {
			order.FailOrder()
			s.ordersRepository.SaveOrder(c, order)
			return err
		}
	}

	order, err = s.ordersRepository.SaveOrder(c, order)
	if err != nil {
		return err
	}

	_, err = s.orderItemsRepository.SaveOrderItems(c, orderItems)
	if err != nil {
		return err
	}

	return nil
}
