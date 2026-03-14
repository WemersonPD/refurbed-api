package models

import (
	"cmp"
	"net/http"
	"slices"
)

type ProductSortBy string

const (
	SortByPriceAsc   ProductSortBy = "price_asc"
	SortByPriceDesc  ProductSortBy = "price_desc"
	SortByBestseller ProductSortBy = "bestseller"
)

// SortProducts sorts the given list of products based on the sort criteria and returns the sorted list.
// In a real application, we would apply the sorting in the database query (ORDER BY clause) instead of in-memory.
func (s ProductSortBy) SortProducts(products []*Product) []*Product {
	switch s {
	case SortByPriceAsc:
		slices.SortFunc(products, func(a, b *Product) int {
			return cmp.Compare(a.DiscountedPrice, b.DiscountedPrice)
		})

	case SortByPriceDesc:
		slices.SortFunc(products, func(a, b *Product) int {
			return cmp.Compare(b.DiscountedPrice, a.DiscountedPrice)
		})

	case SortByBestseller:
		slices.SortFunc(products, func(a, b *Product) int {
			if a.Bestseller != b.Bestseller {
				if a.Bestseller {
					// If a is bestseller and b is not, a should come first (return -1)
					return -1
				}

				// If b is bestseller and a is not, b should come first (return 1)
				return 1
			}

			// If both have the same bestseller status, we keep their original order (return 0).
			return 0
		})
	}

	return products
}

func NewProductSortFromRequest(r *http.Request) ProductSortBy {
	if v := r.URL.Query().Get("sortBy"); v != "" {
		return ProductSortBy(v)
	}

	return SortByBestseller
}
