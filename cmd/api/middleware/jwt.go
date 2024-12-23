package middleware

import (
	"net/http"

	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authentication(app internal.Application) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.CustomAccessJwtClaims)
		},
		SigningKey: []byte(app.Config.Jwt.Access),
	})
}

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := helpers.ContextGetUser(ctx)
		if !user.IsAdmin {
			return &echo.HTTPError{
				Code:    http.StatusForbidden,
				Message: "user has insufficient privileges",
			}
		}

		return next(ctx)
	}
}
