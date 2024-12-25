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

func (s *OrderService) CreateOrder(userId string, req *dto.CreateOrderRequest) (dto.CreateOrderResponse, error) {
	var totalAmount float64
	order := models.Order{
		UserID: userId,
		Status: models.OrderStatusPending,
	}

	tx, err := s.app.Repositories.BeginTransaction()
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}
	defer tx.Rollback()

	for _, orderProduct := range req.Products {
		product, err := s.app.Repositories.Product().UpdateProductStock(tx, orderProduct.ProductID, orderProduct.Quantity)
		if err != nil {
			return dto.CreateOrderResponse{}, err
		}

		totalAmount += product.Price * float64(orderProduct.Quantity)
		order.Products = append(order.Products, models.OrderProduct{
			ProductID: orderProduct.ProductID,
			Quantity:  orderProduct.Quantity,
		})
	}

	order.TotalAmount = totalAmount

	err = s.app.Repositories.Order().CreateOrder(tx, &order)
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	err = tx.Commit()
	if err != nil {
		return dto.CreateOrderResponse{}, err
	}

	return dto.CreateOrderResponse{
		ID:          order.ID,
		UserID:      order.UserID,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
	}, nil
}

func (s *OrderService) UpdateOrderStatus(id string, req *dto.UpdateOrderStatusRequest) (dto.UpdateOrderStatusResponse, error) {
	order, err := s.app.Repositories.Order().GetOrderByID(id)
	if err != nil {
		return dto.UpdateOrderStatusResponse{}, err
	}

	if !models.IsValidOrderStatus(req.Status) {
		return dto.UpdateOrderStatusResponse{}, internal.WrapErrorMessage(internal.ErrBadRequest, "invalid order status")
	}
	status := models.OrderStatus(req.Status)

	err = s.app.Repositories.Order().UpdateOrderStatus(order.ID, status)
	if err != nil {
		return dto.UpdateOrderStatusResponse{}, err
	}

	return dto.UpdateOrderStatusResponse{
		ID:     order.ID,
		Status: string(order.Status),
	}, nil
}

func (s *OrderService) GetOrders(userID string) ([]dto.GetOrderResponse, error) {
	orders, err := s.app.Repositories.Order().GetUserOrders(userID)
	if err != nil {
		return nil, err
	}

	var orderResponses []dto.GetOrderResponse
	for _, order := range orders {
		orderResponses = append(orderResponses, dto.GetOrderResponse{
			ID:          order.ID,
			UserID:      order.UserID,
			Status:      string(order.Status),
			TotalAmount: order.TotalAmount,
			Products:    order.Products,
		})
	}

	return orderResponses, nil
}

func (s *OrderService) CancelOrder(id, userID string) error {
	order, err := s.app.Repositories.Order().GetOrderByID(id)
	if err != nil {
		return err
	}

	if order.Status != models.OrderStatusPending {
		return internal.WrapErrorMessage(internal.ErrBadRequest, "only pending orders can be cancelled")
	}

	if order.UserID != userID {

		return internal.WrapErrorMessage(internal.ErrForbidden, "you can only cancel your own orders")
	}

	return s.app.Repositories.Order().UpdateOrderStatus(order.ID, models.OrderStatusCancelled)
}
