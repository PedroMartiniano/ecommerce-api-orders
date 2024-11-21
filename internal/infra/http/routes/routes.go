package routes

import (
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func InitRoutes(server *gin.Engine) error {
	var err error
	configs.SwaggerConfigure(docs.SwaggerInfo)

	orderRouter := server.Group("/orders")
	err = orderRoutes(orderRouter)
	if err != nil {
		return err
	}

	server.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return nil
}
