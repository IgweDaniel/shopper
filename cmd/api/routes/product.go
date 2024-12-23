package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func registerProductRoutes(e *echo.Echo, handler *handlers.ProductHandler) {
	e.POST("/products", handler.CreateProduct)
	e.GET("/products", handler.GetProducts)
	e.PUT("/products/:id", handler.UpdateProduct)
	e.DELETE("/products/:id", handler.DeleteProduct)
}
