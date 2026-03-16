package models_test

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/utils/testutil"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestProductFilters_IsEmpty(t *testing.T) {
	t.Run("Empty filters", func(t *testing.T) {
		filters := &models.ProductFilters{}

		assert.True(t, filters.IsEmpty())
	})

	t.Run("Non-empty filters with search", func(t *testing.T) {
		filters := &models.ProductFilters{Search: "iphone"}

		assert.False(t, filters.IsEmpty())
	})
}

func TestProductFilters_ApplyProductFilters(t *testing.T) {
	t.Run("Empty filters returns all products", func(t *testing.T) {
		filters := &models.ProductFilters{}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 3)
	})

	t.Run("Filter by search", func(t *testing.T) {
		filters := &models.ProductFilters{Search: "iphone"}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p1", result[0].ID)
	})

	t.Run("Filter by search is case insensitive", func(t *testing.T) {
		filters := &models.ProductFilters{Search: "MACBOOK"}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p3", result[0].ID)
	})

	t.Run("Filter by color", func(t *testing.T) {
		filters := &models.ProductFilters{Color: "blue"}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p1", result[0].ID)
	})

	t.Run("Filter by bestseller true", func(t *testing.T) {
		filters := &models.ProductFilters{Bestseller: testutil.BoolPtr(true)}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 2)
		assert.Equal(t, "p1", result[0].ID)
		assert.Equal(t, "p3", result[1].ID)
	})

	t.Run("Filter by bestseller false", func(t *testing.T) {
		filters := &models.ProductFilters{Bestseller: testutil.BoolPtr(false)}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p2", result[0].ID)
	})

	t.Run("Filter by price range", func(t *testing.T) {
		filters := &models.ProductFilters{
			MinPrice: testutil.Float64Ptr(600.0),
			MaxPrice: testutil.Float64Ptr(800.0),
		}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 2)
		assert.Equal(t, "p1", result[0].ID)
		assert.Equal(t, "p2", result[1].ID)
	})

	t.Run("Filter by category", func(t *testing.T) {
		filters := &models.ProductFilters{Categories: []string{"laptops"}}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p3", result[0].ID)
	})

	t.Run("Filter by multiple categories", func(t *testing.T) {
		filters := &models.ProductFilters{Categories: []string{"smartphones", "laptops"}}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 3)
	})

	t.Run("Filter by brand", func(t *testing.T) {
		filters := &models.ProductFilters{Brands: []string{"samsung"}}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p2", result[0].ID)
	})

	t.Run("Filter by condition", func(t *testing.T) {
		filters := &models.ProductFilters{Conditions: []string{"new"}}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p2", result[0].ID)
	})

	t.Run("Combining multiple filters", func(t *testing.T) {
		filters := &models.ProductFilters{
			Bestseller: testutil.BoolPtr(true),
			Brands:     []string{"apple"},
			Categories: []string{"smartphones"},
		}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Len(t, result, 1)
		assert.Equal(t, "p1", result[0].ID)
	})

	t.Run("No products match filters", func(t *testing.T) {
		filters := &models.ProductFilters{Search: "pixel"}
		products := testutil.SampleProducts()

		result := filters.ApplyProductFilters(products)

		assert.Nil(t, result)
	})
}

func TestNewProductFiltersFromRequest(t *testing.T) {
	t.Run("Parses all query params", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?search=iphone&color=blue&bestseller=true&minPrice=100&maxPrice=500&category=smartphones&brand=apple&condition=new", nil)

		filters := models.NewProductFiltersFromRequest(req)

		assert.Equal(t, "iphone", filters.Search)
		assert.Equal(t, "blue", filters.Color)
		assert.Equal(t, testutil.BoolPtr(true), filters.Bestseller)
		assert.Equal(t, testutil.Float64Ptr(100), filters.MinPrice)
		assert.Equal(t, testutil.Float64Ptr(500), filters.MaxPrice)
		assert.Equal(t, []string{"smartphones"}, filters.Categories)
		assert.Equal(t, []string{"apple"}, filters.Brands)
		assert.Equal(t, []string{"new"}, filters.Conditions)
	})

	t.Run("Empty query returns empty filters", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products", nil)

		filters := models.NewProductFiltersFromRequest(req)

		assert.True(t, filters.IsEmpty())
	})

	t.Run("Price range requires both min and max", func(t *testing.T) {
		req, _ := http.NewRequest("GET", "/products?minPrice=100", nil)

		filters := models.NewProductFiltersFromRequest(req)

		assert.Nil(t, filters.MinPrice)
		assert.Nil(t, filters.MaxPrice)
	})
}
