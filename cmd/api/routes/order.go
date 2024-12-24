package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

func registerOrderRoutes(e *echo.Echo, handler *handlers.OrderHandler) {
	orders := e.Group("orders")
	orders.POST("", handler.CreateOrder, middleware.ValidateDTO(&dto.CreateOrderRequest{}))
	// FIXME: handler to get orders only admin can see and handler to get orders for authenticated user
	orders.GET("/me", handler.GetOrders)
	orders.PUT("/:id", handler.UpdateOrder, middleware.ValidateDTO(&dto.UpdateOrderRequest{}))
	orders.DELETE("/:id", handler.CancelOrder)
}
