package helpers

import (
	"errors"
	"time"

	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/models"
	"github.com/golang-jwt/jwt/v5"
)

const (
	AccessTokenDuration  = 15 * time.Minute
	RefreshTokenDuration = 1 * 24 * time.Hour
)

type CustomAccessJwtClaims struct {
	Id      string `json:"id,omitempty"`
	IsAdmin bool   `json:"adr,omitempty"`
	jwt.RegisteredClaims
}

type CustomRefreshJwtClaims struct {
	Id string `json:"id,omitempty"`
	jwt.RegisteredClaims
}

// GenerateTokens generates the signed access token, refresh token
// and expiration time of the access token.
// Different access and refresh token secrets are used to prevent
// the signed tokens from being used in place of each other.
func GenerateTokens(app *internal.Application, user *models.User) (string, string, int64, error) {
	// generate access token with a lifetime of 15 minutes
	accessTokenExpiration := time.Now().Add(AccessTokenDuration)
	accessClaims := CustomAccessJwtClaims{
		Id:      user.ID,
		IsAdmin: user.IsAdmin,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessTokenExpiration),
		},
	}

	// generate refresh token with a lifetime of 1 day
	refreshClaims := CustomRefreshJwtClaims{
		Id: user.ID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(RefreshTokenDuration)),
		},
	}

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString([]byte(app.Config.Jwt.Access))
	if err != nil {
		return "", "", 0, err
	}

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(app.Config.Jwt.Refresh))
	if err != nil {
		return "", "", 0, err
	}

	return accessToken, refreshToken, accessTokenExpiration.Unix(), err
}

// ValidateRefreshToken validates the provided JWT.
// An error is returned if the token is invalid or expired.
func ValidateRefreshToken(jwtSecret string, signedToken string) (*CustomRefreshJwtClaims, error) {
	// attempt to parse token
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomRefreshJwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecret), nil
		},
	)
	if err != nil {
		return nil, errors.New("invalid or expired token")
	}

	// extract claims from token
	claims, ok := token.Claims.(*CustomRefreshJwtClaims)
	if !ok {
		return nil, errors.New("invalid or expired token")
	}

	return claims, nil
}
