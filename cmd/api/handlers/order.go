package handlers

import (
	"net/http"

	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

type OrderHandler struct {
	Service contracts.OrderService
}

func (h *OrderHandler) CreateOrder(c echo.Context) error {

	req := c.Get("validatedDTO").(*dto.CreateOrderRequest)
	// FIXME: create order for a user
	resp, err := h.Service.CreateOrder(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}

	return c.JSON(http.StatusCreated, resp)
}

func (h *OrderHandler) GetOrders(c echo.Context) error {
	authUser := helpers.ContextGetUser(c)
	orders, err := h.Service.GetOrders(authUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve orders"})
	}

	return c.JSON(http.StatusOK, orders)
}

func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	req := c.Get("validatedDTO").(*dto.UpdateOrderRequest)

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
