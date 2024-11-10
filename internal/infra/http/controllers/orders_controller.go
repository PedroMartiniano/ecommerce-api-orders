package controllers

import (
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/application/services"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/middlewares"
	"github.com/gin-gonic/gin"
)

type OrdersController struct {
	ordersService *services.OrdersService
}

func NewOrdersController(ordersService *services.OrdersService) *OrdersController {
	return &OrdersController{
		ordersService: ordersService,
	}
}

func (o *OrdersController) CreateOrderHandler(c *gin.Context) {
	var request CreateOrderRequest

	if err := c.BindJSON(&request); err != nil {
		sendError(c, http.StatusBadRequest, err.Error())
		return
	}

	user, exists := c.Get("user")
	if !exists {
		sendError(c, http.StatusUnauthorized, "User not found")
		return
	}

	token := c.GetHeader("Authorization")

	userParsed, ok := user.(middlewares.User)
	if !ok {
		sendError(c, http.StatusUnauthorized, "User not found")
		return
	}

	items := []dtos.OrderItemDTO{}
	for _, item := range request.Items {
		items = append(items, dtos.OrderItemDTO{
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
		})
	}
	err := o.ordersService.CreateOrderExecute(c, dtos.OrderDTO{
		Token:       token,
		UserID:      userParsed.ID,
		AddressID:   request.AddressID,
		TotalAmount: request.TotalAmount,
		Items:       items,
		PaymentDetails: dtos.PaymentDetails{
			CardHolder:     request.PaymentDetails.CardHolder,
			CardNumber:     request.PaymentDetails.CardNumber,
			ExpirationDate: request.PaymentDetails.ExpirationDate,
			CVV:            request.PaymentDetails.CVV,
		},
	})
	if err != nil {
		code, message := httpError(err)
		sendError(c, code, message)
		return
	}

	sendSuccess(c, http.StatusCreated, "Order processed successfully")
}
