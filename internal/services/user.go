package services

import (
	"errors"

	"github.com/IgweDaniel/shopper/cmd/api/helpers"
	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"
	"github.com/IgweDaniel/shopper/internal/models"
)

type UserService struct {
	app *internal.Application
}

func NewUserService(app *internal.Application) contracts.UserService {
	return &UserService{app}
}

func (s *UserService) RegisterUser(req *dto.RegisterUserRequest) (dto.RegisterUserResponse, error) {

	PasswordHash, err := helpers.HashPassword(req.Password)
	if err != nil {
		return dto.RegisterUserResponse{}, internal.WrapErrorMessage(err, "failed to hash password")
	}
	user := models.User{
		Email:        req.Email,
		PasswordHash: PasswordHash, // In a real application, make sure to hash the password
	}

	err = s.app.Repositories.User().CreateUser(&user)
	if err != nil {
		if errors.Is(err, internal.ErrDuplicatedKey) {
			return dto.RegisterUserResponse{}, internal.WrapErrorMessage(err, "email already exists")
		}
		return dto.RegisterUserResponse{}, err
	}

	return dto.RegisterUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (s *UserService) LoginUser(req *dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user, err := s.app.Repositories.User().GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, internal.ErrNotFound) {
			return dto.LoginUserResponse{}, internal.WrapErrorMessage(internal.ErrNotAuthorized, "invalid credentials")
		}
		return dto.LoginUserResponse{}, err
	}

	err = helpers.MatchPassword(user.PasswordHash, req.Password)
	if err != nil {
		return dto.LoginUserResponse{}, internal.WrapErrorMessage(internal.ErrNotAuthorized, "invalid credentials")
	}

	accessToken, refreshToken, expiration, err := helpers.GenerateTokens(s.app, user)
	if err != nil {
		return dto.LoginUserResponse{}, internal.WrapErrorMessage(err, "failed to issue tokens")
	}

	return dto.LoginUserResponse{
		AccessToken:          accessToken,
		RefreshToken:         refreshToken,
		AccessTokenExpiresAt: expiration,
	}, nil
}
