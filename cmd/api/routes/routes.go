package routes

import (
	"net/http"

	"github.com/IgweDaniel/shopper/cmd/api/handlers"
	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/database"

	"github.com/labstack/echo/v4"

	_ "github.com/IgweDaniel/shopper/docs"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type Router struct {
	App  *internal.Application
	Echo *echo.Echo
}

func NewRouter(app *internal.Application, e *echo.Echo) *Router {
	return &Router{
		App:  app,
		Echo: e,
	}
}

func RegisterRoutes(app *internal.Application, db database.Service, services *contracts.Services) http.Handler {
	e := echo.New()
	var whiteList = []string{"*"}

	e.Validator = helpers.NewCustomValidator()
	e.HideBanner = true
	e.Debug = true

	e.Use(echo.MiddlewareFunc(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Before(func() {
				origin := c.Request().Header.Get("Origin")
				for _, allowedOrigins := range whiteList {
					// we are in development
					if allowedOrigins == "*" {
						c.Response().Header().Set("Access-Control-Allow-Origin", "*")
						break
					}
					if allowedOrigins == origin {
						c.Response().Header().Set("Access-Control-Allow-Origin", origin)
					}
				}

				c.Response().Header().Set("Access-Control-Allow-Headers", "*")
				c.Response().Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, PATCH, OPTIONS")
				c.Response().Header().Add("Vary", "Origin")

			})
			return next(c)
		}
	}))

	e.Pre(middleware.RemoveTrailingSlash())

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	router := NewRouter(app, e)

	router.registerUserRoutes(handlers.NewUserHandler(services.User))
	router.registerOrderRoutes(handlers.NewOrderHandler(services.Order))
	router.registerProductRoutes(handlers.NewProductHandler(services.Product))

	e.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, db.Health())
	})

	e.GET("/swagger/*", echoSwagger.WrapHandler)
	return e
}
