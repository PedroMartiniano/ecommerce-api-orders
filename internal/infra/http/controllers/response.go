package controllers

import (
	"errors"
	"net/http"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"
	"github.com/gin-gonic/gin"
)

var logger = configs.GetLogger()

func httpError(err error) (code int, message string) {
	var myErr configs.Error

	if errors.As(err, &myErr) {
		message = myErr.AppError().Error()

		switch myErr.TypeError() {
		case configs.ErrInternalServer:
			code = http.StatusInternalServerError
		case configs.ErrNotFound:
			code = http.StatusNotFound
		case configs.ErrBadRequest:
			code = http.StatusBadRequest
		}
	}

	return
}

func sendError(c *gin.Context, code int, err string) {
	logger.Error(err)
	c.JSON(code, gin.H{"success": false, "message": err})
}

func sendSuccess(c *gin.Context, code int, data any) {
	c.JSON(code, gin.H{"success": true, "data": data})
}

type errorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

type orderResponse1 struct {
	Success bool         `json:"success"`
	Data    dtos.OrderResDTO `json:"data"`
}

