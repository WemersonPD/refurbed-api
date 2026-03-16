package testutil

import (
	"assignment-backend/pkg/models"
	"context"

	"github.com/stretchr/testify/mock"
)

type MockProductsRepository struct {
	mock.Mock
}

func (m *MockProductsRepository) ListProducts(ctx context.Context, filters *models.ProductFilters, sortBy models.ProductSortBy, pagination *models.Pagination) ([]*models.Product, error) {
	args := m.Called(ctx, filters, sortBy, pagination)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]*models.Product), args.Error(1)
}

func (m *MockProductsRepository) CountProducts(ctx context.Context, filters *models.ProductFilters) (int, error) {
	args := m.Called(ctx, filters)
	return args.Int(0), args.Error(1)
}
