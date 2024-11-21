package main

import (
	"fmt"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/routes"
	"github.com/gin-gonic/gin"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	configs.Init()

	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	err := routes.InitRoutes(server)
	if err != nil {
		panic(err.Error())
	}

	port := configs.GetEnv("PORT")
	server.Run(fmt.Sprintf(":%s", port))
}
