package models

import (
	"fmt"
	"net/http"
	"strconv"
)

type Pagination struct {
	Limit  int
	Offset int
}

// ApplyProductsPagination applies the pagination to the given list of products and returns the paginated result.
// In a real application, we would apply the pagination in the database query (e.g., using LIMIT and OFFSET in SQL).
func (p *Pagination) ApplyProductsPagination(products []*Product) []*Product {
	if p.Offset >= len(products) {
		return []*Product{}
	}
	// Keeping the end index within the bounds of the products slice.
	end := min(p.Offset+p.Limit, len(products))

	return products[p.Offset:end]
}

func NewPagination(r *http.Request) (*Pagination, error) {
	q := r.URL.Query()

	limitParameter := q.Get("limit")
	if limitParameter == "" {
		return nil, fmt.Errorf("limit query parameter is required")
	}

	offsetParameter := q.Get("offset")
	if offsetParameter == "" {
		return nil, fmt.Errorf("offset query parameter is required")
	}

	limit, err := strconv.Atoi(limitParameter)
	if err != nil {
		return nil, fmt.Errorf("invalid limit query parameter")
	}

	offset, err := strconv.Atoi(offsetParameter)
	if err != nil {
		return nil, fmt.Errorf("invalid offset query parameter")
	}

	return &Pagination{
		Limit:  limit,
		Offset: offset,
	}, nil
}
