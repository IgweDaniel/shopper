package services

import (
	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"
	"github.com/IgweDaniel/shopper/internal/models"
)

type OrderService struct {
	app *internal.Application
}

func NewOrderService(app *internal.Application) contracts.OrderService {
	return &OrderService{app}
}

func (s *OrderService) CreateOrder(req *dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	order := models.Order{
		ProductID: req.ProductID,
		Quantity:  req.Quantity,
		Status:    models.OrderStatusPending,
	}

	err := s.app.Repositories.Order.CreateOrder(&order)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	return dto.CreateOrderResponse{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Status:    string(order.Status),
	}, nil
}

func (s *OrderService) UpdateOrder(id string, req *dto.UpdateOrderRequest) (dto.UpdateOrderResponse, error) {
	order, err := s.app.Repositories.Order.GetOrderByID(id)
	if err != nil {
		return dto.UpdateOrderResponse{}, err
	}

	order.Status = models.OrderStatus(req.Status)

	err = s.app.Repositories.Order.UpdateOrder(order)
	if err != nil {
		return dto.UpdateOrderResponse{}, err
	}

	return dto.UpdateOrderResponse{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Status:    string(order.Status),
	}, nil
}

func (s *OrderService) GetOrderByID(id string) (dto.GetOrderResponse, error) {
	order, err := s.app.Repositories.Order.GetOrderByID(id)
	if err != nil {
		return dto.GetOrderResponse{}, err
	}

	return dto.GetOrderResponse{
		ID:        order.ID,
		ProductID: order.ProductID,
		Quantity:  order.Quantity,
		Status:    string(order.Status),
	}, nil
}

func (s *OrderService) GetOrders(userID string) ([]dto.GetOrderResponse, error) {
	orders, err := s.app.Repositories.Order.GetUserOrders(userID)
	if err != nil {
		return nil, err
	}

	var orderResponses []dto.GetOrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, dto.GetOrderResponse{
			ID:        order.ID,
			ProductID: order.ProductID,
			Quantity:  order.Quantity,
			Status:    string(order.Status),
		})
	}

	return orderResponses, nil
}

func (s *OrderService) CancelOrder(id string) error {
	order, err := s.app.Repositories.Order.GetOrderByID(id)
	if err != nil {
		return err
	}

	order.Status = models.OrderStatusCancelled

	return s.app.Repositories.Order.UpdateOrder(order)
}
