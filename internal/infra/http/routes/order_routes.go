package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/adapters"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/controllers"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/middlewares"
	"github.com/gin-gonic/gin"
)

func orderRoutes(router *gin.RouterGroup) {
	orderService := adapters.NewOrderServiceAdapter(configs.DB)
	ordersController := controllers.NewOrdersController(orderService)

	router.POST("/", middlewares.VerifyToken, ordersController.CreateOrderHandler)
}
