package repository

import (
	"database/sql"
	"errors"

	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/models"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) contracts.UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (email, password, is_admin) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, user.Email, user.Password, user.IsAdmin).Scan(&user.ID)
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, email, password, is_admin FROM users WHERE email = $1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.Password, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("user not found")
		}
		return nil, err
	}
	return user, nil
}
