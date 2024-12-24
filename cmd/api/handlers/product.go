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

func (h *ProductHandler) CreateProduct(c echo.Context) error {

	req := c.Get("validatedDTO").(*dto.CreateProductRequest)

	resp, err := h.Service.CreateProduct(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *ProductHandler) GetProducts(c echo.Context) error {
	products, err := h.Service.GetProducts()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve products"})
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHandler) UpdateProduct(c echo.Context) error {
	id := c.Param("id")

	req := c.Get("validatedDTO").(*dto.UpdateProductRequest)

	resp, err := h.Service.UpdateProduct(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update product"})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *ProductHandler) DeleteProduct(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.DeleteProduct(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete product"})
	}

	return c.NoContent(http.StatusNoContent)
}
