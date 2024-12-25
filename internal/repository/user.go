package repository

import (
	"database/sql"
	"strings"

	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/models"
)

type PostgresUserRepository struct {
	DB *sql.DB
}

const duplicateEmail = "users_"

func NewPostgresUserRepository(db *sql.DB) contracts.UserRepository {
	return &PostgresUserRepository{DB: db}
}

func (r *PostgresUserRepository) CreateUser(user *models.User) error {
	query := "INSERT INTO users (email, password_hash, is_admin) VALUES ($1, $2, $3) RETURNING id"
	err := r.DB.QueryRow(query, user.Email, user.PasswordHash, user.IsAdmin).Scan(&user.ID)
	if err != nil {
		switch {
		case
			strings.Contains(err.Error(), duplicateEmail):
			return internal.ErrDuplicatedKey
		default:
			return err
		}
	}
	return nil
}

func (r *PostgresUserRepository) GetUserByEmail(email string) (*models.User, error) {
	user := &models.User{}
	query := "SELECT id, email, password_hash, is_admin FROM users WHERE email = $1"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.IsAdmin)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.ErrNotFound
		}
		return nil, err
	}
	return user, nil
}
