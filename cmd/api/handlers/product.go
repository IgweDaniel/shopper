package handlers

import (
	"net/http"

	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

type ProductHandler struct {
	Service contracts.ProductService
}

func NewProductHandler(service contracts.ProductService) *ProductHandler {
	return &ProductHandler{Service: service}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with name, description, and price
// @Tags products
// @Accept json
// @Produce json
// @Param product body dto.CreateProductRequest true "Product creation request"
// @Success 201 {object} helpers.ApiResponse{data=dto.CreateProductResponse}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 401 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 403 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /products [post]
func (h *ProductHandler) CreateProduct(c echo.Context) error {

	req := c.Get("validatedDTO").(*dto.CreateProductRequest)

	resp, err := h.Service.CreateProduct(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}

// GetProducts godoc
// @Summary Get all products
// @Description Get all products
// @Tags products
// @Produce json
// @Success 200 {object} helpers.ApiResponse{data=[]dto.GetProductResponse}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /products [get]
func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.Service.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve products"})
	}

	return c.JSON(http.StatusOK, products)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product's name, description, and price
// @Tags products
// @Accept json
// @Produce json
// @Param id path string true "Product ID"
// @Param product body dto.UpdateProductRequest true "Product update request"
// @Success 200 {object} helpers.ApiResponse{data=dto.UpdateProductResponse}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 401 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 403 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /products/{id} [put]
func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	req := c.Get("validatedDTO").(*dto.UpdateProductRequest)

	resp, err := h.Service.UpdateProduct(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by ID
// @Tags products
// @Produce json
// @Param id path string true "Product ID"
// @Success 204
// @Success 201 {object}  helpers.ApiResponse{}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /products/{id} [delete]
func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	return c.NoContent(http.StatusNoContent)
}
