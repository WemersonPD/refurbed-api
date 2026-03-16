package services_test

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/services"
	utils_cache "assignment-backend/pkg/utils/cache"
	"assignment-backend/pkg/utils/testutil"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupService() (services.ProductService, *testutil.MockProductsRepository) {
	mockRepo := &testutil.MockProductsRepository{}
	cache := utils_cache.NewCache[*models.ProductsResponse](0)
	svc := services.NewProductService(mockRepo, cache)
	return svc, mockRepo
}

func TestGetProducts(t *testing.T) {
	t.Run("Success returns products and count", func(t *testing.T) {
		svc, mockRepo := setupService()
		products := testutil.SampleProducts()
		filters := &models.ProductFilters{}
		sort := models.SortByBestseller
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		mockRepo.On("ListProducts", mock.Anything, filters, sort, pagination).Return(products, nil).Once()
		mockRepo.On("CountProducts", mock.Anything, filters).Return(3, nil).Once()

		resp, err := svc.GetProducts(context.TODO(), "all-products", filters, sort, pagination)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Len(t, resp.Products, 3)
		assert.Equal(t, 3, resp.Count)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Cache hit skips repository call", func(t *testing.T) {
		svc, mockRepo := setupService()
		products := testutil.SampleProducts()
		filters := &models.ProductFilters{}
		sort := models.SortByBestseller
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		// First call populates cache.
		mockRepo.On("ListProducts", mock.Anything, filters, sort, pagination).Return(products, nil).Once()
		mockRepo.On("CountProducts", mock.Anything, filters).Return(3, nil).Once()

		resp1, err := svc.GetProducts(context.TODO(), "cached-key", filters, sort, pagination)
		assert.NoError(t, err)
		assert.NotNil(t, resp1)

		// Second call with same key should return cached result without calling repo again.
		resp2, err := svc.GetProducts(context.TODO(), "cached-key", filters, sort, pagination)
		assert.NoError(t, err)
		assert.NotNil(t, resp2)
		assert.Equal(t, resp1.Count, resp2.Count)
		assert.Equal(t, len(resp1.Products), len(resp2.Products))

		// Repository should only have been called once.
		mockRepo.AssertExpectations(t)
	})

	t.Run("ListProducts error returns error", func(t *testing.T) {
		svc, mockRepo := setupService()
		filters := &models.ProductFilters{}
		sort := models.SortByBestseller
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		mockRepo.On("ListProducts", mock.Anything, filters, sort, pagination).Return(nil, fmt.Errorf("database error")).Once()

		resp, err := svc.GetProducts(context.TODO(), "error-key", filters, sort, pagination)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "database error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("CountProducts error returns error", func(t *testing.T) {
		svc, mockRepo := setupService()
		products := testutil.SampleProducts()
		filters := &models.ProductFilters{}
		sort := models.SortByBestseller
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		mockRepo.On("ListProducts", mock.Anything, filters, sort, pagination).Return(products, nil).Once()
		mockRepo.On("CountProducts", mock.Anything, filters).Return(0, fmt.Errorf("count error")).Once()

		resp, err := svc.GetProducts(context.TODO(), "count-error-key", filters, sort, pagination)

		assert.Error(t, err)
		assert.Nil(t, resp)
		assert.Equal(t, "count error", err.Error())
		mockRepo.AssertExpectations(t)
	})

	t.Run("Empty results returns empty response", func(t *testing.T) {
		svc, mockRepo := setupService()
		filters := &models.ProductFilters{Search: "nonexistent"}
		sort := models.SortByBestseller
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		mockRepo.On("ListProducts", mock.Anything, filters, sort, pagination).Return([]*models.Product{}, nil).Once()
		mockRepo.On("CountProducts", mock.Anything, filters).Return(0, nil).Once()

		resp, err := svc.GetProducts(context.TODO(), "empty-key", filters, sort, pagination)

		assert.NoError(t, err)
		assert.NotNil(t, resp)
		assert.Empty(t, resp.Products)
		assert.Equal(t, 0, resp.Count)
		mockRepo.AssertExpectations(t)
	})
}
