package contracts

import "github.com/IgweDaniel/shopper/internal/dto"

type Services struct {
	User    UserService
	Product ProductService
	Order   OrderService
}

type UserService interface {
	RegisterUser(req *dto.RegisterUserRequest) (dto.RegisterUserResponse, error)
	LoginUser(req *dto.LoginUserRequest) (dto.LoginUserResponse, error)
}

type ProductService interface {
	CreateProduct(req *dto.CreateProductRequest) (dto.CreateProductResponse, error)
	UpdateProduct(id string, req *dto.UpdateProductRequest) (dto.UpdateProductResponse, error)
	GetProductByID(id string) (dto.GetProductResponse, error)
	GetProducts() ([]dto.GetProductResponse, error)
	DeleteProduct(id string) error
}

type OrderService interface {
	CreateOrder(req *dto.CreateOrderRequest) (dto.CreateOrderResponse, error)
	UpdateOrder(id string, req *dto.UpdateOrderRequest) (dto.UpdateOrderResponse, error)
	GetOrderByID(id string) (dto.GetOrderResponse, error)
	GetOrders(userId string) ([]dto.GetOrderResponse, error)
	CancelOrder(id string) error
}
