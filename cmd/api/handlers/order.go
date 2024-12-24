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

func NewOrderHandler(service contracts.OrderService) *OrderHandler {
	return &OrderHandler{Service: service}
}

// GetOrders godoc
// @Summary Get all orders
// @Description Get a list of all orders
// @Tags orders
// @Produce json
// @Success 200 {object} helpers.ApiResponse{data=dto.GetOrderResponse}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /orders [get]
func (h *OrderHandler) CreateOrder(c echo.Context) error {

	req := c.Get("validatedDTO").(*dto.CreateOrderRequest)
	// FIXME: create order for a user
	resp, err := h.Service.CreateOrder(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create order"})
	}

	return c.JSON(http.StatusCreated, resp)
}

// GetOrder godoc
// @Summary Get an order by ID
// @Description Get details of an order by ID
// @Tags orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 200 {object} helpers.ApiResponse{data=dto.GetOrderResponse}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 404 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /orders/{id} [get]
func (h *OrderHandler) GetOrders(c echo.Context) error {
	authUser := helpers.ContextGetUser(c)
	orders, err := h.Service.GetOrders(authUser.ID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to retrieve orders"})
	}

	return c.JSON(http.StatusOK, orders)
}

// UpdateOrder godoc
// @Summary Update an order
// @Description Update an order's details
// @Tags orders
// @Accept json
// @Produce json
// @Param id path string true "Order ID"
// @Param order body dto.UpdateOrderRequest true "Order update request"
// @Success 200 {object} helpers.ApiResponse{data=dto.UpdateOrderResponse}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 401 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 403 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /orders/{id} [put]
func (h *OrderHandler) UpdateOrder(c echo.Context) error {
	id := c.Param("id")
	req := c.Get("validatedDTO").(*dto.UpdateOrderRequest)

	resp, err := h.Service.UpdateOrder(id, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update order"})
	}

	return c.JSON(http.StatusOK, resp)
}

// DeleteOrder godoc
// @Summary Delete an order
// @Description Delete an order by ID
// @Tags orders
// @Produce json
// @Param id path string true "Order ID"
// @Success 204 {object} helpers.ApiResponse{}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 401 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 403 {object} helpers.ApiResponse{message=string,success=bool}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /orders/{id} [delete]
func (h *OrderHandler) CancelOrder(c echo.Context) error {
	id := c.Param("id")
	if err := h.Service.CancelOrder(id); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to cancel order"})
	}

	return c.NoContent(http.StatusNoContent)
}
