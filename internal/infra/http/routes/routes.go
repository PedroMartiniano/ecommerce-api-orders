package routes

import "github.com/gin-gonic/gin"

func InitRoutes(server *gin.Engine) {
	orderRouter := server.Group("/orders")
	orderRoutes(orderRouter)
}
