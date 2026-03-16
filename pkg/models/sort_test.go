package models_test

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/utils/testutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductSortBy_SortProducts(t *testing.T) {
	t.Run("Sort by price ascending", func(t *testing.T) {
		products := testutil.SampleProducts()

		result := models.SortByPriceAsc.SortProducts(products)

		assert.Equal(t, "p2", result[0].ID) // 629.99
		assert.Equal(t, "p1", result[1].ID) // 749.99
		assert.Equal(t, "p3", result[2].ID) // 1199.99
	})

	t.Run("Sort by price descending", func(t *testing.T) {
		products := testutil.SampleProducts()

		result := models.SortByPriceDesc.SortProducts(products)

		assert.Equal(t, "p3", result[0].ID) // 1199.99
		assert.Equal(t, "p1", result[1].ID) // 749.99
		assert.Equal(t, "p2", result[2].ID) // 629.99
	})

	t.Run("Sort by bestseller", func(t *testing.T) {
		products := testutil.SampleProducts()

		result := models.SortByBestseller.SortProducts(products)

		// p1 and p3 are bestsellers, p2 is not
		assert.True(t, result[0].Bestseller)
		assert.True(t, result[1].Bestseller)
		assert.False(t, result[2].Bestseller)
	})
}

func TestNewProductSortFromRequest(t *testing.T) {
	t.Run("Returns sortBy from query param", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?sortBy=price_asc", nil)

		result := models.NewProductSortFromRequest(req)

		assert.Equal(t, models.SortByPriceAsc, result)
	})

	t.Run("Defaults to bestseller when no sortBy", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products", nil)

		result := models.NewProductSortFromRequest(req)

		assert.Equal(t, models.SortByBestseller, result)
	})
}
