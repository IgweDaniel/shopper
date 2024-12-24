package handlers

import (
	"net/http"

	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	Service contracts.UserService
}

func NewUserHandler(service contracts.UserService) *UserHandler {
	return &UserHandler{Service: service}
}

// Register godoc
// @Summary Register a new user
// @Description Register a new user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.RegisterUserRequest true "User registration request"
// @Success 201 {object}  helpers.ApiResponse{data=dto.RegisterUserResponse}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /user [post]
func (h *UserHandler) Register(c echo.Context) error {
	req := c.Get("validatedDTO").(*dto.RegisterUserRequest)
	resp, err := h.Service.RegisterUser(req)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return c.JSON(http.StatusCreated, helpers.BuildResponse("account created", resp))
}

// Login godoc
// @Summary Login a user
// @Description Login a user with email and password
// @Tags users
// @Accept json
// @Produce json
// @Param user body dto.LoginUserRequest true "User login request"
// @Success 201 {object}  helpers.ApiResponse{data=dto.LoginUserResponse}
// @Failure 400 {object} helpers.ApiResponse{message=string,success=bool,data=map[string]string}
// @Failure 500 {object} helpers.ApiResponse{message=string,success=bool}
// @Router /user/auth [post]
func (h *UserHandler) Login(c echo.Context) error {
	req := c.Get("validatedDTO").(*dto.LoginUserRequest)

	resp, err := h.Service.LoginUser(req)
	if err != nil {
		return helpers.HandleError(c, err)
	}

	return c.JSON(http.StatusOK, helpers.BuildResponse("auth successful", resp))
}
