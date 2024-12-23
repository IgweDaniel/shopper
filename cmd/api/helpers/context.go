package helpers

import (
	"github.com/IgweDaniel/shopper/internal/models"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type ContextKey string

const ContextKeyUser = "user"

// ContextGetUser returns the current user data stored in the request context.
// Only the ID and Address of a user are stored.
func ContextGetUser(ctx echo.Context) models.User {
	token, ok := ctx.Get(ContextKeyUser).(*jwt.Token)
	if !ok {
		return models.User{}
	}
	claims, ok := token.Claims.(*CustomAccessJwtClaims)
	if !ok {
		return models.User{}
	}

	return models.User{
		ID:      claims.ID,
		IsAdmin: claims.IsAdmin,
	}
}
