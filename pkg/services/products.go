package services

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/repositories"
	"context"
)

type ProductService interface {
	GetProducts(ctx context.Context, filters *models.ProductFilters, sort models.ProductSortBy) ([]*models.Product, error)
}

type productService struct {
	productsRepository repositories.ProductsRepository
}

func NewProductService() ProductService {
	return &productService{
		productsRepository: repositories.NewProductsRepository(),
	}
}

func (s *productService) GetProducts(ctx context.Context, filters *models.ProductFilters, sort models.ProductSortBy) ([]*models.Product, error) {
	// Get the data from cache.

	// Get the data from database if not in cache.
	products, err := s.productsRepository.GetProducts(ctx, filters, sort)

	return products, err
}
