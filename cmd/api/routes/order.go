package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func registerOrderRoutes(e *echo.Echo, handler *handlers.OrderHandler) {
	e.POST("/orders", handler.CreateOrder)
	e.GET("/orders", handler.GetOrders)
	e.PUT("/orders/:id", handler.UpdateOrder)
	e.DELETE("/orders/:id", handler.CancelOrder)
}
