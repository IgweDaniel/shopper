package services

import (
	"errors"

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
	user := models.User{
		Email:    req.Email,
		Password: req.Password, // In a real application, make sure to hash the password
	}

	err := s.app.Repositories.User.CreateUser(&user)
	if err != nil {
		return dto.RegisterUserResponse{}, err
	}

	return dto.RegisterUserResponse{
		ID:    user.ID,
		Email: user.Email,
	}, nil
}

func (s *UserService) LoginUser(req *dto.LoginUserRequest) (dto.LoginUserResponse, error) {
	user, err := s.app.Repositories.User.GetUserByEmail(req.Email)
	if err != nil {
		return dto.LoginUserResponse{}, err
	}

	// In a real application, make sure to compare the hashed password
	if user.Password != req.Password {
		return dto.LoginUserResponse{}, errors.New("invalid credentials")
	}

	// Generate a token (this is a placeholder, implement your own token generation logic)
	token := "generated-jwt-token"

	return dto.LoginUserResponse{
		Token: token,
	}, nil
}
