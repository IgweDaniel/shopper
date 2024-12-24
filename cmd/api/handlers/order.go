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

// CreateOrder godoc
//
//	@Summary		Create a new order
//	@Description	Create a new order
//	@Tags			orders
//	@Security		JWT
//	@Accept			json
//	@Produce		json
//	@Param			order	body		dto.CreateOrderRequest	true	"Order create request"
//	@Success		201		{object}	helpers.ApiResponse{data=dto.CreateOrderResponse}
//	@Failure		500		{object}	helpers.ApiResponse{message=string,success=bool}
//	@Router			/orders [post]
func (h *OrderHandler) CreateOrder(c echo.Context) error {
	userID := helpers.ContextGetUser(c).ID
	req := c.Get("validatedDTO").(*dto.CreateOrderRequest)

	resp, err := h.Service.CreateOrder(userID, req)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return c.JSON(http.StatusCreated, resp)
}

// GetUserOrders godoc
//
//	@Summary		Get all orders for authenticated user
//	@Description	Get all orders for authenticated user
//	@Tags			orders
//	@Security		JWT
//	@Produce		json
//	@Success		200		{object}	helpers.ApiResponse{data=[]dto.GetOrderResponse}
//	@Failure		500		{object}	helpers.ApiResponse{message=string,success=bool}
//	@Router			/orders/me [get]
func (h *OrderHandler) GetOrders(c echo.Context) error {
	userID := helpers.ContextGetUser(c).ID
	orders, err := h.Service.GetOrders(userID)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, helpers.BuildResponse("orders fetched", orders))
}

// UpdateOrderStatus godoc
//
//	@Summary		Update order status
//	@Description	Update the status of an order by ID
//	@Tags			admin
//	@Produce		json
//	@Security		JWT
//	@Param			id		path	string					true	"Order ID"
//	@Param			status	body	dto.UpdateOrderStatusRequest	true	"Order status update request"
//	@Success		200		{object}	helpers.ApiResponse{data=dto.UpdateOrderStatusResponse}
//	@Failure		400		{object}	helpers.ApiResponse{message=string,success=bool}
//	@Failure		401		{object}	helpers.ApiResponse{message=string,success=bool}
//	@Failure		403		{object}	helpers.ApiResponse{message=string,success=bool}
//	@Failure		500		{object}	helpers.ApiResponse{message=string,success=bool}
//	@Router			/orders/{id}/status [put]
func (h *OrderHandler) UpdateOrderStatus(c echo.Context) error {
	id := c.Param("id")
	req := c.Get("validatedDTO").(*dto.UpdateOrderStatusRequest)

	resp, err := h.Service.UpdateOrderStatus(id, req)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, helpers.BuildResponse("order status updated", resp))
}

// CancelOrder godoc
//
//	@Summary		Cancel an order
//	@Description	Cancel an order by ID
//	@Tags			orders
//	@Produce		json
//	@Security		JWT
//	@Param			id	path	string	true	"Order ID"
//	@Success		204	{object}	helpers.ApiResponse{}
//	@Failure		400	{object}	helpers.ApiResponse{message=string,success=bool}
//	@Failure		401	{object}	helpers.ApiResponse{message=string,success=bool}
//	@Failure		403	{object}	helpers.ApiResponse{message=string,success=bool}
//	@Failure		500	{object}	helpers.ApiResponse{message=string,success=bool}
//	@Router			/orders/{id}/cancel [put]
func (h *OrderHandler) CancelOrder(c echo.Context) error {
	id := c.Param("id")
	userID := helpers.ContextGetUser(c).ID
	if err := h.Service.CancelOrder(id, userID); err != nil {
		return helpers.HandleError(c, err)
	}

	return c.NoContent(http.StatusNoContent)
}
