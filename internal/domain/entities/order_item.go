package entities

import (
	"time"

	"github.com/PedroMartiniano/ecommerce-api-orders/internal/domain/vo"
)

type OrderItem struct {
	ID         vo.UUID     `json:"id"`
	OrderID    vo.UUID     `json:"order_id"`
	ProductID  vo.UUID     `json:"product_id"`
	Quantity   vo.Quantity `json:"quantity"`
	UnitPrice  vo.Amount   `json:"price"`
	TotalPrice vo.Amount   `json:"total"`
	CreatedAt  time.Time   `json:"created_at"`
}

func CreateNewOrderItem(orderID string, productID string, quantity int, unitPrice float64, totalPrice float64) (OrderItem, error) {
	quantityParsed, err := vo.NewQuantity(quantity)
	if err != nil {
		return OrderItem{}, err
	}

	unitPriceParsed, err := vo.NewAmount(unitPrice)
	if err != nil {
		return OrderItem{}, err
	}

	totalPriceParsed, err := vo.NewAmount(totalPrice)
	if err != nil {
		return OrderItem{}, err
	}

	return OrderItem{
		ID:         vo.NewUUID(""),
		OrderID:    vo.NewUUID(orderID),
		ProductID:  vo.NewUUID(productID),
		Quantity:   quantityParsed,
		UnitPrice:  unitPriceParsed,
		TotalPrice: totalPriceParsed,
		CreatedAt:  time.Now(),
	}, nil
}

func InstantiateOrderItem(id string, orderID string, productID string, quantity int, unitPrice float64, totalPrice float64, createdAt time.Time) (OrderItem, error) {
	quantityParsed, err := vo.NewQuantity(quantity)
	if err != nil {
		return OrderItem{}, err
	}

	unitPriceParsed, err := vo.NewAmount(unitPrice)
	if err != nil {
		return OrderItem{}, err
	}

	totalPriceParsed, err := vo.NewAmount(totalPrice)
	if err != nil {
		return OrderItem{}, err
	}

	return OrderItem{
		ID:         vo.NewUUID(id),
		OrderID:    vo.NewUUID(orderID),
		ProductID:  vo.NewUUID(productID),
		Quantity:   quantityParsed,
		UnitPrice:  unitPriceParsed,
		TotalPrice: totalPriceParsed,
		CreatedAt:  createdAt,
	}, nil
}

func (o *OrderItem) GetID() string {
	return o.ID.GetValue()
}

func (o *OrderItem) GetOrderID() string {
	return o.OrderID.GetValue()
}

func (o *OrderItem) GetProductID() string {
	return o.ProductID.GetValue()
}

func (o *OrderItem) GetQuantity() int {
	return o.Quantity.GetValue()
}

func (o *OrderItem) GetUnitPrice() float64 {
	return o.UnitPrice.GetValue()
}

func (o *OrderItem) GetTotalPrice() float64 {
	return o.TotalPrice.GetValue()
}

func (o *OrderItem) GetCreatedAt() time.Time {
	return o.CreatedAt
}
