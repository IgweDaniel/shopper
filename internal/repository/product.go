package repository

import (
	"database/sql"
	"errors"

	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/models"
)

type PostgresProductRepository struct {
	DB *sql.DB
}

func NewPostgresProductRepository(db *sql.DB) contracts.ProductRepository {
	return &PostgresProductRepository{DB: db}
}

func (r *PostgresProductRepository) CreateProduct(product *models.Product) error {
	query := "INSERT INTO products (name, description, price) VALUES ($1, $2, $3) RETURNING id"
	return r.DB.QueryRow(query, product.Name, product.Description, product.Price).Scan(&product.ID)
}

func (r *PostgresProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, name, description, price FROM products"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *PostgresProductRepository) GetProductByID(id string) (*models.Product, error) {
	product := &models.Product{}
	query := "SELECT id, name, description, price FROM products WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("product not found")
		}
		return nil, err
	}
	return product, nil
}

func (r *PostgresProductRepository) UpdateProduct(product *models.Product) error {
	query := "UPDATE products SET name = $1, description = $2, price = $3 WHERE id = $4"
	_, err := r.DB.Exec(query, product.Name, product.Description, product.Price, product.ID)
	return err
}

func (r *PostgresProductRepository) DeleteProduct(id string) error {
	query := "DELETE FROM products WHERE id = $1"
	_, err := r.DB.Exec(query, id)
	return err
}
