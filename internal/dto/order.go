package dto

import "github.com/IgweDaniel/shopper/internal/models"

type (
	CreateOrderRequest struct {
		Products []OrderProduct `json:"products" validate:"required,dive"`
	}

	OrderProduct struct {
		ProductID string `json:"product_id" validate:"required,uuid"`
		Quantity  int    `json:"quantity" validate:"required,gt=0"`
	}

	CreateOrderResponse struct {
		ID          string  `json:"id"`
		UserID      string  `json:"user_id"`
		Status      string  `json:"status"`
		TotalAmount float64 `json:"total_amount"`
	}
)

type GetOrderResponse struct {
	ID          string                `json:"id"`
	UserID      string                `json:"user_id"`
	Status      string                `json:"status"`
	TotalAmount float64               `json:"total_amount"`
	Products    []models.OrderProduct `json:"products"`
}

type (
	UpdateOrderStatusRequest struct {
		Status string `json:"status" validate:"required"`
	}

	UpdateOrderStatusResponse struct {
		ID     string `json:"id"`
		Status string `json:"status"`
	}
)
