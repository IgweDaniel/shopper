package handlers

import (
	"net/http"

	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Service contracts.OrderService
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {
	req := new(dto.CreateOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp, err := h.Service.CreateOrder(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *OrderHandler) GetOrders(c echo.Context) error {
	userID := c.QueryParam("user_id")
	orders, err := h.Service.GetOrders(userID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve orders"})
	}

	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	req := new(dto.UpdateOrderRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	resp, err := h.Service.UpdateOrder(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order"})
	}

	return c.JSON(http.StatusOK, resp)
}

func (h *OrderHandler) CancelOrder(c echo.Context) error {
	id := c.Param("id")

	if err := h.Service.CancelOrder(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to cancel order"})
	}

	return c.NoContent(http.StatusNoContent)
}
