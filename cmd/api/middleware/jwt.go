package middleware

import (
	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func Authentication(app *internal.Application) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(helpers.CustomAccessJwtClaims)
		},
		SigningKey: []byte(app.Config.Jwt.Access),
		ErrorHandler: func(c echo.Context, err error) error {
			return helpers.HandleError(c, internal.WrapErrorMessage(internal.ErrNotAuthorized, "unauthorized"))

		},
	})
}

func RequireAdmin(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		user := helpers.ContextGetUser(ctx)

		if !user.IsAdmin {
			return helpers.HandleError(ctx, internal.WrapErrorMessage(internal.ErrForbidden, "admin privilege required"))
		}

		return next(ctx)
	}
}
