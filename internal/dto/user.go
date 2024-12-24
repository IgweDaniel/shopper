package dto

type (
	RegisterUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,min=6"`
	}

	RegisterUserResponse struct {
		ID    string `json:"id"`
		Email string `json:"email"`
	}
)

type (
	LoginUserRequest struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}

	LoginUserResponse struct {
		AccessToken          string `json:"access_token"`
		RefreshToken         string `json:"refresh_token"`
		AccessTokenExpiresAt int64  `json:"aceess_token_expires_at"`
	}
)
