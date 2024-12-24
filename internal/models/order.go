package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCancelled OrderStatus = "cancelled" // Fixed value
	OrderStatusShipped   OrderStatus = "shipped"
)

type Order struct {
	ID          string         `json:"id"`
	UserID      string         `json:"user_id"`
	Status      OrderStatus    `json:"status"`
	Products    []OrderProduct `json:"products"`
	TotalAmount float64        `json:"total_amount"`
}

type OrderProduct struct {
	ProductID string `json:"product_id"`
	Quantity  int    `json:"quantity"`
}

func IsValidOrderStatus(status string) bool {
	switch OrderStatus(status) {
	case OrderStatusPending, OrderStatusCancelled, OrderStatusShipped:
		return true
	default:
		return false
	}
}
