package ports

import "github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"

type IProductGateway interface {
	GetProductByID(productID string) (dtos.Product, error)
	UpdateProductQuantity(dto dtos.UpdateProductQuantityDTO) error
}
