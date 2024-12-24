package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"
)

func (r *Router) registerProductRoutes(handler *handlers.ProductHandler) {

	products := r.Echo.Group("products")
	products.GET("", handler.GetProducts, middleware.Authentication(r.App))
	products.GET("/:id", handler.GetProduct, middleware.Authentication(r.App))
	products.POST("", handler.CreateProduct, middleware.Authentication(r.App), middleware.RequireAdmin, middleware.ValidateDTO(&dto.CreateProductRequest{}))
	products.PUT("/:id", handler.UpdateProduct, middleware.Authentication(r.App), middleware.RequireAdmin, middleware.ValidateDTO(&dto.UpdateProductRequest{}))
	products.DELETE("/:id", handler.DeleteProduct, middleware.Authentication(r.App), middleware.RequireAdmin)
}
