package controllers_test

import (
	"assignment-backend/internal/controllers"
	"assignment-backend/pkg/models"
	utils_response "assignment-backend/pkg/utils/response"
	"assignment-backend/pkg/utils/testutil"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type paginatedResponse = utils_response.ApiResponse[utils_response.Pagination[[]*models.Product]]
type errorResponse = utils_response.ApiResponse[map[string]string]

func setupController() (controllers.ProductsController, *testutil.MockProductService) {
	mockService := &testutil.MockProductService{}
	ctrl := controllers.NewProductsController(mockService)
	return ctrl, mockService
}

func TestGetProducts(t *testing.T) {
	t.Run("Success returns paginated products", func(t *testing.T) {
		ctrl, mockService := setupController()
		products := testutil.SampleProducts()

		mockService.On("GetProducts", mock.Anything, "limit=10&offset=0", mock.Anything, models.SortByBestseller, &models.Pagination{Limit: 10, Offset: 0}).
			Return(&models.ProductsResponse{Products: products, Count: 3}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/products?limit=10&offset=0", nil)
		rec := httptest.NewRecorder()

		ctrl.GetProducts(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp paginatedResponse
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.True(t, resp.OK)
		assert.Len(t, resp.Data.Data, 3)
		assert.Equal(t, 3, resp.Data.Total)
		assert.Equal(t, 10, resp.Data.Limit)
		assert.Equal(t, 0, resp.Data.Offset)
		mockService.AssertExpectations(t)
	})

	t.Run("Missing limit returns 400", func(t *testing.T) {
		ctrl, _ := setupController()

		req := httptest.NewRequest(http.MethodGet, "/products?offset=0", nil)
		rec := httptest.NewRecorder()

		ctrl.GetProducts(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp errorResponse
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.False(t, resp.OK)
		assert.Contains(t, resp.Data["error"], "limit")
	})

	t.Run("Missing offset returns 400", func(t *testing.T) {
		ctrl, _ := setupController()

		req := httptest.NewRequest(http.MethodGet, "/products?limit=10", nil)
		rec := httptest.NewRecorder()

		ctrl.GetProducts(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

		var resp errorResponse
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.False(t, resp.OK)
		assert.Contains(t, resp.Data["error"], "offset")
	})

	t.Run("Service error returns 500", func(t *testing.T) {
		ctrl, mockService := setupController()

		mockService.On("GetProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(nil, fmt.Errorf("service error")).Once()

		req := httptest.NewRequest(http.MethodGet, "/products?limit=10&offset=0", nil)
		rec := httptest.NewRecorder()

		ctrl.GetProducts(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

		var resp errorResponse
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.False(t, resp.OK)
		assert.Equal(t, "Failed to retrieve products", resp.Data["error"])
		mockService.AssertExpectations(t)
	})

	t.Run("Empty results returns 200 with empty data", func(t *testing.T) {
		ctrl, mockService := setupController()

		mockService.On("GetProducts", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything).
			Return(&models.ProductsResponse{Products: []*models.Product{}, Count: 0}, nil).Once()

		req := httptest.NewRequest(http.MethodGet, "/products?limit=10&offset=0&search=nonexistent", nil)
		rec := httptest.NewRecorder()

		ctrl.GetProducts(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

		var resp paginatedResponse
		err := json.NewDecoder(rec.Body).Decode(&resp)
		assert.NoError(t, err)
		assert.True(t, resp.OK)
		assert.Empty(t, resp.Data.Data)
		assert.Equal(t, 0, resp.Data.Total)
		mockService.AssertExpectations(t)
	})
}
