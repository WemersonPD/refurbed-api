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
