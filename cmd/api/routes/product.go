package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

func registerProductRoutes(e *echo.Echo, handler *handlers.ProductHandler) {

	products := e.Group("products")
	products.POST("", handler.CreateProduct, middleware.ValidateDTO(&dto.CreateProductRequest{}))
	products.GET("", handler.GetProducts)
	products.PUT("/:id", handler.UpdateProduct, middleware.ValidateDTO(&dto.UpdateProductRequest{}))
	products.DELETE("/:id", handler.DeleteProduct)
}
