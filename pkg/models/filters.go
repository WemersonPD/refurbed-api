package models

import (
	"net/http"
	"strconv"
)

type ProductFilters struct {
	Search     string
	Color      string
	Bestseller *bool
	MinPrice   *float64
	MaxPrice   *float64
}

func (f *ProductFilters) IsEmpty() bool {
	return f.Search == "" && f.Color == "" && f.Bestseller == nil && f.MinPrice == nil && f.MaxPrice == nil
}

// ApplyProductFilters applies the filters to the given list of products and returns the filtered list.
// In a real application, we would apply these filters in the database query (WHERE clause) instead of in-memory.
func (f *ProductFilters) ApplyProductFilters(products []*Product) (filteredProducts []*Product) {
	if f.IsEmpty() {
		return products
	}

	for _, product := range products {
		if f.Search != "" && !product.Name.Contains(f.Search) {
			continue
		}

		if f.Color != "" && !product.Colors.Contains(f.Color) {
			continue
		}

		if f.Bestseller != nil && product.Bestseller != *f.Bestseller {
			continue
		}

		if f.MinPrice != nil && f.MaxPrice != nil {
			if product.DiscountedPrice < *f.MinPrice || product.DiscountedPrice > *f.MaxPrice {
				continue
			}
		}

		filteredProducts = append(filteredProducts, product)
	}

	return filteredProducts
}

func NewProductFiltersFromRequest(r *http.Request) *ProductFilters {
	q := r.URL.Query()
	f := &ProductFilters{
		Search: q.Get("search"),
		Color:  q.Get("color"),
	}

	if v := q.Get("bestseller"); v != "" {
		b, _ := strconv.ParseBool(v)
		f.Bestseller = &b
	}

	// Min and Max are a range filter, so we only set them if both are provided
	minPrice := q.Get("minPrice")
	maxPrice := q.Get("maxPrice")
	if minPrice != "" && maxPrice != "" {
		minPriceParsed, _ := strconv.ParseFloat(minPrice, 64)
		maxPriceParsed, _ := strconv.ParseFloat(maxPrice, 64)
		f.MinPrice = &minPriceParsed
		f.MaxPrice = &maxPriceParsed
	}

	return f
}
