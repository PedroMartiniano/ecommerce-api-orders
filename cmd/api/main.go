package main

import (
	"fmt"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/infra/http/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	configs.Init()

	gin.SetMode(gin.DebugMode)

	server := gin.Default()

	routes.InitRoutes(server)

	port := configs.GetEnv("PORT")
	server.Run(fmt.Sprintf(":%s", port))
}
