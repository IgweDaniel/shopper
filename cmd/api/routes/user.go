package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"

	"github.com/labstack/echo/v4"
)

func registerUserRoutes(e *echo.Echo, handler *handlers.UserHandler) {

	users := e.Group("users")
	users.POST("", handler.Register, middleware.ValidateDTO(&dto.RegisterUserRequest{}))
	users.POST("/auth", handler.Login, middleware.ValidateDTO(&dto.LoginUserRequest{}))
}
