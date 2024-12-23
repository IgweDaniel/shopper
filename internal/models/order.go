package models

type OrderStatus string

const (
	OrderStatusPending   OrderStatus = "pending"
	OrderStatusCancelled OrderStatus = "cancelled" // Fixed value
)

type Order struct {
	ID        string      `json:"id"`
	UserID    string      `json:"user_id"`
	ProductID string      `json:"product_id"`
	Quantity  int         `json:"quantity"`
	Status    OrderStatus `json:"status"`
}
