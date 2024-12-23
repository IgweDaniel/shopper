package routes

import (
	"net/http"

	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/database"

	"github.com/labstack/echo/v4"

	_ "github.com/IgweDaniel/shopper/docs"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func RegisterRoutes(db database.Service, services *contracts.Services) http.Handler {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	userHandler := &handlers.UserHandler{Service: services.User}
	orderHandler := &handlers.OrderHandler{Service: services.Order}
	productHandler := &handlers.ProductHandler{Service: services.Product}

	registerUserRoutes(e, userHandler)
	registerOrderRoutes(e, orderHandler)
	registerProductRoutes(e, productHandler)

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, db.Health())
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
