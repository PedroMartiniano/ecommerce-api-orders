package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/adapters"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/controllers"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/middlewares"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/queue"
	"github.com/gin-gonic/gin"
)

func orderRoutes(router *gin.RouterGroup) error {
	var err error

	rabbitMQ := queue.NewRabbitMQQueue()
	err = rabbitMQ.Connect()
	if err != nil {
		return err
	}
	orderService := adapters.NewOrderServiceAdapter(configs.DB, rabbitMQ)
	ordersController := controllers.NewOrdersController(orderService)

	router.POST("/", middlewares.VerifyToken, ordersController.CreateOrderHandler)
	return nil
}
