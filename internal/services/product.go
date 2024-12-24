package services

import (
	"github.com/IgweDaniel/shopper/internal"
	"github.com/IgweDaniel/shopper/internal/contracts"
	"github.com/IgweDaniel/shopper/internal/dto"
	"github.com/IgweDaniel/shopper/internal/models"
)

type ProductService struct {
	app *internal.Application
}

func NewProductService(app *internal.Application) contracts.ProductService {
	return &ProductService{app}
}

func (s *ProductService) CreateProduct(req *dto.CreateProductRequest) (dto.CreateProductResponse, error) {
	product := models.Product{
		Name:        req.Name,
		Description: req.Description,
		Price:       req.Price,
		Stock:       req.Stock,
	}

	err := s.app.Repositories.Product.CreateProduct(&product)
	if err != nil {
		return dto.CreateProductResponse{}, err
	}

	return dto.CreateProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

func (s *ProductService) UpdateProduct(id string, req *dto.UpdateProductRequest) (dto.UpdateProductResponse, error) {
	product, err := s.app.Repositories.Product.UpdateUnderLock(id, func(product *models.Product) error {

		if req.Name != nil {
			product.Name = *req.Name
		}
		if req.Description != nil {
			product.Description = *req.Description
		}
		if req.Price != nil {
			product.Price = *req.Price
		}
		if req.Stock != nil {
			product.Stock = *req.Stock
		}

		return nil
	})
	if err != nil {
		return dto.UpdateProductResponse{}, err
	}

	return dto.UpdateProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
		Stock:       product.Stock,
	}, nil
}

func (s *ProductService) GetProductByID(id string) (dto.GetProductResponse, error) {
	product, err := s.app.Repositories.Product.GetProductByID(id)
	if err != nil {
		return dto.GetProductResponse{}, err
	}

	return dto.GetProductResponse{
		ID:          product.ID,
		Name:        product.Name,
		Description: product.Description,
		Price:       product.Price,
	}, nil
}

func (s *ProductService) GetProducts() ([]dto.GetProductResponse, error) {
	products, err := s.app.Repositories.Product.GetProducts()
	if err != nil {
		return nil, err
	}

	var productResponses []dto.GetProductResponse
	for _, product := range products {
		productResponses = append(productResponses, dto.GetProductResponse{
			ID:          product.ID,
			Name:        product.Name,
			Description: product.Description,
			Price:       product.Price,
		})
	}

	return productResponses, nil
}

func (s *ProductService) DeleteProduct(id string) error {
	return s.app.Repositories.Product.DeleteProduct(id)
}
