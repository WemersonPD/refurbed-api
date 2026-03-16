package testutil

import (
	"assignment-backend/pkg/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) GetProducts(ctx context.Context, cacheKey string, filters *models.ProductFilters, sort models.ProductSortBy, pagination *models.Pagination) (*models.ProductsResponse, error) {
	args := m.Called(ctx, cacheKey, filters, sort, pagination)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.ProductsResponse), args.Error(1)
}
