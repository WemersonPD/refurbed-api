package models

import "net/http"

type ProductSortBy string

const (
	SortByPriceAsc   ProductSortBy = "price_asc"
	SortByPriceDesc  ProductSortBy = "price_desc"
	SortByBestseller ProductSortBy = "bestseller"
)

func NewProductSortFromRequest(r *http.Request) ProductSortBy {

	if v := r.URL.Query().Get("sortBy"); v != "" {
		return ProductSortBy(v)
	}

	return SortByBestseller
}
