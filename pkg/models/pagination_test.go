package models_test

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/utils/testutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPagination_ApplyProductsPagination(t *testing.T) {
	t.Run("First page", func(t *testing.T) {
		pagination := &models.Pagination{Limit: 2, Offset: 0}
		products := testutil.SampleProducts()

		result := pagination.ApplyProductsPagination(products)

		assert.Len(t, result, 2)
		assert.Equal(t, "p1", result[0].ID)
		assert.Equal(t, "p2", result[1].ID)
	})

	t.Run("Second page", func(t *testing.T) {
		pagination := &models.Pagination{Limit: 2, Offset: 2}
		products := testutil.SampleProducts()

		result := pagination.ApplyProductsPagination(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p3", result[0].ID)
	})

	t.Run("Offset beyond products length returns empty", func(t *testing.T) {
		pagination := &models.Pagination{Limit: 2, Offset: 10}
		products := testutil.SampleProducts()

		result := pagination.ApplyProductsPagination(products)

		assert.Empty(t, result)
	})

	t.Run("Empty products list", func(t *testing.T) {
		pagination := &models.Pagination{Limit: 6, Offset: 0}

		result := pagination.ApplyProductsPagination([]*models.Product{})

		assert.Empty(t, result)
	})
}

func TestNewPagination(t *testing.T) {
	t.Run("Parses valid limit and offset", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?limit=6&offset=12", nil)

		pagination, err := models.NewPagination(req)

		assert.NoError(t, err)
		assert.Equal(t, 6, pagination.Limit)
		assert.Equal(t, 12, pagination.Offset)
	})

	t.Run("Missing limit returns error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?offset=0", nil)

		pagination, err := models.NewPagination(req)

		assert.Nil(t, pagination)
		assert.EqualError(t, err, "limit query parameter is required")
	})

	t.Run("Missing offset returns error", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?limit=6", nil)

		pagination, err := models.NewPagination(req)

		assert.Nil(t, pagination)
		assert.EqualError(t, err, "offset query parameter is required")
	})
}
