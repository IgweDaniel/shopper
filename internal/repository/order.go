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
	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	query := "INSERT INTO orders (user_id, status, total_amount) VALUES ($1, $2, $3) RETURNING id"
	err = tx.QueryRow(query, order.UserID, order.Status, order.TotalAmount).Scan(&order.ID)
	if err != nil {
		tx.Rollback()
		return err
	}

	for _, product := range order.Products {
		query = "INSERT INTO order_products (order_id, product_id, quantity) VALUES ($1, $2, $3)"
		_, err = tx.Exec(query, order.ID, product.ProductID, product.Quantity)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	return tx.Commit()
}

func (r *PostgresOrderRepository) GetUserOrders(userID string) ([]models.Order, error) {
	query := `
        SELECT o.id, o.user_id, o.status, o.total_amount, op.product_id, op.quantity
        FROM orders o
        LEFT JOIN order_products op ON o.id = op.order_id
        WHERE o.user_id = $1
    `
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	orders := make(map[string]*models.Order)
	for rows.Next() {
		var orderID, userID, productID string
		var status models.OrderStatus
		var totalAmount float64
		var quantity int

		if err := rows.Scan(&orderID, &userID, &status, &totalAmount, &productID, &quantity); err != nil {
			return nil, err
		}

		if _, exists := orders[orderID]; !exists {
			orders[orderID] = &models.Order{
				ID:          orderID,
				UserID:      userID,
				Status:      status,
				TotalAmount: totalAmount,
				Products:    []models.OrderProduct{},
			}
		}

		orders[orderID].Products = append(orders[orderID].Products, models.OrderProduct{
			ProductID: productID,
			Quantity:  quantity,
		})
	}

	orderList := make([]models.Order, 0, len(orders))
	for _, order := range orders {
		orderList = append(orderList, *order)
	}

	return orderList, nil
}

func (r *PostgresOrderRepository) GetOrderByID(id string) (models.Order, error) {
	query := `
		SELECT o.id, o.user_id, o.status, o.total_amount, op.product_id, op.quantity
		FROM orders o
		LEFT JOIN order_products op ON o.id = op.order_id
		WHERE o.id = $1
	`
	rows, err := r.DB.Query(query, id)
	if err != nil {
		return models.Order{}, err
	}
	defer rows.Close()

	order := models.Order{}
	for rows.Next() {
		var orderID, userID, productID string
		var status models.OrderStatus
		var totalAmount float64
		var quantity int

		if err := rows.Scan(&orderID, &userID, &status, &totalAmount, &productID, &quantity); err != nil {
			return models.Order{}, err
		}

		if order.ID == "" {
			order.ID = orderID
			order.UserID = userID
			order.Status = status
			order.TotalAmount = totalAmount
		}

		order.Products = append(order.Products, models.OrderProduct{
			ProductID: productID,
			Quantity:  quantity,
		})
	}

	if order.ID == "" {
		return models.Order{}, errors.New("order not found")
	}

	return order, nil
}

func (r *PostgresOrderRepository) UpdateOrderStatus(id string, status models.OrderStatus) error {
	query := "UPDATE orders SET status = $1 WHERE id = $2"
	_, err := r.DB.Exec(query, status, id)
	return err
}

func (r *PostgresOrderRepository) DeleteOrder(id string) error {
	query := "DELETE FROM orders WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
