package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"

	"github.com/labstack/echo/v4"
)

func registerUserRoutes(e *echo.Echo, handler *handlers.UserHandler) {
	e.POST("/register", handler.Register)
	e.POST("/login", handler.Login)
}
