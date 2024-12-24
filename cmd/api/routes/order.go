package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"
)

func (r *Router) registerOrderRoutes(handler *handlers.OrderHandler) {
	orders := r.Echo.Group("orders")
	orders.POST("", handler.CreateOrder, middleware.Authentication(r.App), middleware.ValidateDTO(&dto.CreateOrderRequest{}))
	orders.GET("/me", handler.GetOrders, middleware.Authentication(r.App))
	orders.PUT("/:id/cancel", handler.CancelOrder, middleware.Authentication(r.App))
	orders.PUT("/:id/status", handler.UpdateOrderStatus, middleware.RequireAdmin, middleware.Authentication(r.App), middleware.ValidateDTO(&dto.UpdateOrderStatusRequest{}))
}
