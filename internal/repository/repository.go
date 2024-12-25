package repository

import (
	"database/sql"

	"github.com/IgweDaniel/shopper/internal/contracts"
)

type PostgresTransaction struct {
	tx *sql.Tx
}

func (t *PostgresTransaction) Commit() error {
	return t.tx.Commit()
}

func (t *PostgresTransaction) Rollback() error {
	return t.tx.Rollback()
}

func NewPostgresTransaction(tx *sql.Tx) contracts.Transaction {
	return &PostgresTransaction{tx: tx}
}

type PostgresRepository struct {
	DB *sql.DB
}

func (r *PostgresRepository) BeginTransaction() (contracts.Transaction, error) {
	tx, err := r.DB.Begin()
	if err != nil {
		return nil, err
	}
	return NewPostgresTransaction(tx), nil
}

func (r *PostgresRepository) Product() contracts.ProductRepository {
	return &PostgresProductRepository{DB: r.DB}
}

func (r *PostgresRepository) Order() contracts.OrderRepository {
	return &PostgresOrderRepository{DB: r.DB}
}
func (r *PostgresRepository) User() contracts.UserRepository {
	return &PostgresUserRepository{DB: r.DB}
}
