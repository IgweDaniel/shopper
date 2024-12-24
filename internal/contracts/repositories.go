package contracts

import (
	"github.com/IgweDaniel/shopper/internal/models"
)

type Repositories struct {
	User    UserRepository
	Order   OrderRepository
	Product ProductRepository
}

// FIXME: CRUD for each model, add pagination and filtering (order and product)

type UserRepository interface {
	CreateUser(user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetProductByID(id string) (*models.Product, error)
	GetProducts() ([]models.Product, error)
	UpdateProduct(product *models.Product) error
	DeleteProduct(id string) error
	UpdateUnderLock(productID string, updaterFunc func(product *models.Product) error) (*models.Product, error)
	// GetProductForUpdate(tx *sql.Tx, productID string) (*models.Product, error)
}

type OrderRepository interface {
	CreateOrder(order *models.Order) error
	GetOrderByID(id string) (models.Order, error)
	GetUserOrders(userID string) ([]models.Order, error)
	UpdateOrderStatus(id string, status models.OrderStatus) error
	DeleteOrder(id string) error
}
