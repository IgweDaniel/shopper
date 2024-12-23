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
		Token string `json:"token"`
	}
)
