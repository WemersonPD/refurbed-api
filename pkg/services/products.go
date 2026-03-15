package services

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/repositories"
	"context"
)

type ProductService interface {
	CountProducts(ctx context.Context, filters *models.ProductFilters, sort models.ProductSortBy, pagination *models.Pagination) (products []*models.Product, count int, err error)
}

type productService struct {
	productsRepository repositories.ProductsRepository
}

func NewProductService() ProductService {
	return &productService{
		productsRepository: repositories.NewProductsRepository(),
	}
}

func (s *productService) CountProducts(ctx context.Context, filters *models.ProductFilters, sort models.ProductSortBy, pagination *models.Pagination) (products []*models.Product, count int, err error) {
	// TODO: Get the data from cache.

	// Get the data from database if not in cache.
	products, err = s.productsRepository.ListProducts(ctx, filters, sort, pagination)
	if err != nil {
		return nil, 0, err
	}

	count, err = s.productsRepository.CountProducts(ctx, filters)
	if err != nil {
		return nil, 0, err
	}

	// TODO: Save the data to cache.

	return products, count, err
}
