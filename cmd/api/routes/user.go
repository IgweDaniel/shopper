package routes

import (
	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/middleware"
	"github.com/IgweDaniel/shopper/internal/dto"
)

func (r *Router) registerUserRoutes(handler *handlers.UserHandler) {

	users := r.Echo.Group("users")
	users.POST("", handler.Register, middleware.ValidateDTO(&dto.RegisterUserRequest{}))
	users.POST("/auth", handler.Login, middleware.ValidateDTO(&dto.LoginUserRequest{}))
}
