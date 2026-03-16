package services

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/repositories"
	utils_cache "assignment-backend/pkg/utils/cache"
	"context"
)

type ProductService interface {
	GetProducts(ctx context.Context, cacheKey string, filters *models.ProductFilters, sort models.ProductSortBy, pagination *models.Pagination) (response *models.ProductsResponse, err error)
}

type productService struct {
	productsRepository repositories.ProductsRepository
	productsCache      *utils_cache.Cache[*models.ProductsResponse]
}

func NewProductService(repo repositories.ProductsRepository, cache *utils_cache.Cache[*models.ProductsResponse]) ProductService {
	if repo == nil {
		repo = repositories.NewProductsRepository(nil)
	}

	if cache == nil {
		cache = utils_cache.NewCache[*models.ProductsResponse](0)
	}

	return &productService{
		productsRepository: repo,
		productsCache:      cache,
	}
}

func (s *productService) GetProducts(ctx context.Context, cacheKey string, filters *models.ProductFilters, sort models.ProductSortBy, pagination *models.Pagination) (response *models.ProductsResponse, err error) {
	response, hasCache := s.productsCache.Get(cacheKey)
	if hasCache {
		return response, err
	}

	// Get the data from database if not in cache.
	products, err := s.productsRepository.ListProducts(ctx, filters, sort, pagination)
	if err != nil {
		return nil, err
	}

	count, err := s.productsRepository.CountProducts(ctx, filters)
	if err != nil {
		return nil, err
	}

	response = &models.ProductsResponse{
		Products: products,
		Count:    count,
	}

	s.productsCache.Set(cacheKey, response)

	return &models.ProductsResponse{
		Products: products,
		Count:    count,
	}, nil
}
