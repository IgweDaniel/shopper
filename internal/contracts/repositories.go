package contracts

import (
	"github.com/IgweDaniel/shopper/internal/models"
)

type Transaction interface {
	Commit() error
	Rollback() error
}

type Repositories interface {
	BeginTransaction() (Transaction, error)
	Product() ProductRepository
	Order() OrderRepository
	User() UserRepository
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
	UpdateProductStock(tx Transaction, productID string, quantity int) (*models.Product, error)
	Update(id string, updates map[string]interface{}) error
}

type OrderRepository interface {
	CreateOrder(tx Transaction, order *models.Order) error
	GetOrderByID(id string) (models.Order, error)
	GetUserOrders(userID string) ([]models.Order, error)
	UpdateOrderStatus(id string, status models.OrderStatus) error
	DeleteOrder(id string) error
}
