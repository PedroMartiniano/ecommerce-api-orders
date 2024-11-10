package gateways

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	pg "github.com/PedroMartiniano/ecommerce-api-orders/internal/application/ports/gateways"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/configs"
	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/dtos"
)

var logger = configs.GetLogger()

type ProductGateway struct {
}

func NewProductGateway() pg.IProductGateway {
	return &ProductGateway{}
}

type response struct {
	Success bool         `json:"success"`
	Data    dtos.Product `json:"data"`
}

func (p *ProductGateway) GetProductByID(productID string) (dtos.Product, error) {
	base_url := configs.GetEnv("PRODUCT_SERVICE_URL")
	url := fmt.Sprintf("%s/products/%s", base_url, productID)

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return dtos.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return dtos.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return dtos.Product{}, configs.NewError(configs.ErrInternalServer, errors.New("product not found"))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return dtos.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}

	var response response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return dtos.Product{}, configs.NewError(configs.ErrInternalServer, err)
	}

	return response.Data, nil
}

func (p *ProductGateway) UpdateProductQuantity(dto dtos.UpdateProductQuantityDTO) error {
	base_url := configs.GetEnv("PRODUCT_SERVICE_URL")
	url := fmt.Sprintf("%s/products/%s/stock", base_url, dto.ProductID)

	body := map[string]interface{}{
		"quantity":  dto.Quantity,
		"operation": dto.Operation,
	}

	jsonBody, err := json.Marshal(body)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	req, err := http.NewRequest("PUT", url, bytes.NewBuffer(jsonBody))
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", dto.Token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return configs.NewError(configs.ErrInternalServer, err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Errorf("Failed to update product quantity: %v", resp.StatusCode)
		return configs.NewError(configs.ErrInternalServer, errors.New("failed to update product quantity"))
	}

	return nil
}
