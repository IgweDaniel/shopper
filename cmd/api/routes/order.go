package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"
)

func (r *Router) registerOrderRoutes(handler *handlers.OrderHandler) {
	orders := r.Echo.Group("orders")
	orders.POST("", handler.CreateOrder, middleware.Authentication(r.App), middleware.ValidateDTO(&dto.CreateOrderRequest{}))
	// FIXME: handler to get orders only admin can see and handler to get orders for authenticated user
	orders.GET("/me", handler.GetOrders, middleware.Authentication(r.App))
	// 	FIXME: this should handle canceling or changing order state by admin
	orders.PUT("/:id", handler.UpdateOrder, middleware.Authentication(r.App), middleware.ValidateDTO(&dto.UpdateOrderRequest{}))
	// TODO: handler to cancel an order
	orders.DELETE("/:id", handler.CancelOrder)
}
