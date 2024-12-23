package repository

import (
	"database/sql"
	"errors"

	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/models"
)

type PostgresOrderRepository struct {
	DB *sql.DB
}

func NewPostgresOrderRepository(db *sql.DB) contracts.OrderRepository {
	return &PostgresOrderRepository{DB: db}
}

func (r *PostgresOrderRepository) CreateOrder(order *models.Order) error {
	query := "INSERT INTO orders (user_id, product_id, quantity, status) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.DB.QueryRow(query, order.UserID, order.ProductID, order.Quantity, order.Status).Scan(&order.ID)
}

func (r *PostgresOrderRepository) GetUserOrders(userID string) ([]models.Order, error) {
	query := "SELECT id, user_id, product_id, quantity, status FROM orders WHERE user_id = $1"
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		if err := rows.Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status); err != nil {
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (r *PostgresOrderRepository) GetOrderByID(id string) (*models.Order, error) {
	order := &models.Order{}
	query := "SELECT id, user_id, product_id, quantity, status FROM orders WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&order.ID, &order.UserID, &order.ProductID, &order.Quantity, &order.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("order not found")
		}
		return nil, err
	}
	return order, nil
}

func (r *PostgresOrderRepository) UpdateOrder(order *models.Order) error {
	query := "UPDATE orders SET user_id = $1, product_id = $2, quantity = $3, status = $4 WHERE id = $5"
	_, err := r.DB.Exec(query, order.UserID, order.ProductID, order.Quantity, order.Status, order.ID)
	return err
}

func (r *PostgresOrderRepository) DeleteOrder(id string) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
