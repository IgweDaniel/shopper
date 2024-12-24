package handlers

import (
	"net/http"

	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service contracts.UserService
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterUserRequest true "User registration request"
// @Success 201 {object} dto.RegisterUserResponse
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /user [post]
func (h *UserHandler) Register(c echo.Context) error {
	req := c.Get("validatedDTO").(*dto.RegisterUserRequest)
	resp, err := h.Service.RegisterUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusCreated, resp)
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.LoginUserRequest true "User login request"
// @Success 200 {object} dto.LoginUserResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Router /user/auth [post]
func (h *UserHandler) Login(c echo.Context) error {
	req := c.Get("validatedDTO").(*dto.LoginUserRequest)

	resp, err := h.Service.LoginUser(req)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
	}

	return c.JSON(http.StatusOK, resp)
}
