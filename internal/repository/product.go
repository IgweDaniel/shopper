package repository

import (
	"database/sql"

	"github.com/IgweDaniel/shopper/internal"
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
	query := "INSERT INTO products (name, description, price, stock) VALUES ($1, $2, $3, $4) RETURNING id"
	return r.DB.QueryRow(query, product.Name, product.Description, product.Price, product.Stock).Scan(&product.ID)
}

func (r *PostgresProductRepository) GetProducts() ([]models.Product, error) {
	query := "SELECT id, name, description, price, stock FROM products"
	rows, err := r.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	products := []models.Product{}
	for rows.Next() {
		product := models.Product{}
		if err := rows.Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock); err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func (r *PostgresProductRepository) GetProductByID(id string) (*models.Product, error) {
	product := &models.Product{}
	query := "SELECT id, name, description, price, stock FROM products WHERE id = $1"
	err := r.DB.QueryRow(query, id).Scan(&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, internal.WrapErrorMessage(internal.ErrNotFound, "product not found")
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

func (r *PostgresProductRepository) UpdateProductStock(tx contracts.Transaction, productID string, quantity int) (*models.Product, error) {
	pgTx := tx.(*PostgresTransaction).tx
	var product models.Product

	err := pgTx.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id = $1 FOR UPDATE", productID).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return &product, internal.WrapErrorMessage(internal.ErrNotFound, "product not found")
		}
		return &product, err
	}

	if product.Stock < quantity {
		return &product, internal.WrapErrorMessage(internal.ErrBadRequest, "insufficient stock")
	}

	product.Stock -= quantity

	_, err = pgTx.Exec("UPDATE products SET stock = $1 WHERE id = $2", product.Stock, product.ID)
	return &product, err
}

func (r *PostgresProductRepository) Update(id string, updates map[string]interface{}) error {

	tx, err := r.DB.Begin()
	if err != nil {
		return err
	}

	defer tx.Rollback()

	var product models.Product

	err = tx.QueryRow("SELECT id, name, description, price, stock FROM products WHERE id = $1 FOR UPDATE", id).Scan(
		&product.ID, &product.Name, &product.Description, &product.Price, &product.Stock)
	if err != nil {
		if err == sql.ErrNoRows {
			return internal.WrapErrorMessage(internal.ErrNotFound, "product not found")
		}
		return err
	}

	for key, value := range updates {
		switch key {
		case "name":
			product.Name = value.(string)
		case "description":
			product.Description = value.(string)
		case "price":
			product.Price = value.(float64)
		case "stock":
			product.Stock = value.(int)
		}
	}

	_, err = tx.Exec("UPDATE products SET name = $1, description = $2, price = $3, stock= $4 WHERE id = $5", product.Name, product.Description, product.Price, product.Stock, product.ID)
	if err != nil {
		return err
	}

	return tx.Commit()
}
