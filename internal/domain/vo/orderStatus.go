package vo

type Status string

const (
	OrderStatusPending   Status = "pending"
	OrderStatusConfirmed Status = "confirmed"
	OrderStatusCancelled Status = "cancelled"
	OrderStatusShipped   Status = "shipped"
	OrderStatusDelivered Status = "delivered"
)

type OrderStatus struct {
	value Status
}

func NewOrderStatus(value Status) OrderStatus {

	
	return OrderStatus{
		value: value,
	}
}


func (o OrderStatus) GetValue() Status {
	return o.value
}