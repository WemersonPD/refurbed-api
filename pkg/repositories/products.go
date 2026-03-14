package repositories

import (
	"assignment-backend/pkg/models"
	utils_jsongloader "assignment-backend/pkg/utils/jsongloader"
	utils_numbers "assignment-backend/pkg/utils/numbers"
	"cmp"
	"context"
	"slices"
)

const (
	METADATA_FILE_PATH = "data/metadata.json"
	DETAILS_FILE_PATH  = "data/details.json"
)

type ProductsRepository interface {
	// GetProducts reads the metadata and details, merges them and returns the list of products.
	GetProducts(ctx context.Context, filters *models.ProductFilters, sortBy models.ProductSortBy) (products []*models.Product, err error)
}

type productsRepository struct {
	jsonLoader utils_jsongloader.JSONLoader
}

func NewProductsRepository() ProductsRepository {
	return &productsRepository{
		jsonLoader: utils_jsongloader.NewJSONLoader(),
	}
}

func (p *productsRepository) getProductsMetadata() (products []*models.ProductMetadata, err error) {
	err = p.jsonLoader.LoadJSONFile(METADATA_FILE_PATH, &products)

	return products, err
}

func (p *productsRepository) getProductsDetails() (details []*models.ProductDetail, err error) {
	err = p.jsonLoader.LoadJSONFile(DETAILS_FILE_PATH, &details)

	return details, err
}

// joinProductsMetadataAndDetails merges the metadata and details slices into a slice of Product.
// If we had a database, this would be a JOIN query instead of in-memory merging.
func (p *productsRepository) joinProductsMetadataAndDetails(metadata []*models.ProductMetadata, details []*models.ProductDetail) (products []*models.Product) {
	// Merge metadata and details by ID
	detailsMap := make(map[string]*models.ProductDetail)
	for _, detail := range details {
		detailsMap[detail.ID] = detail
	}

	for _, meta := range metadata {
		if detail, exists := detailsMap[meta.ID]; exists {
			// If we were using a database, we would calculate the discounted price in the QUERY.
			discountedPrice := utils_numbers.CalculateDiscountedPrice(meta.BasePrice, detail.DiscountPercent)

			product := &models.Product{
				// Detail fields
				DiscountPercent: detail.DiscountPercent,
				Bestseller:      detail.Bestseller,
				Colors:          detail.Colors,
				ImageURL:        meta.ImageURL,
				Stock:           detail.Stock,

				// Metadata fields
				ID:        meta.ID,
				Name:      meta.Name,
				BasePrice: meta.BasePrice,

				// // Calculated fields
				DiscountedPrice: discountedPrice,
			}
			products = append(products, product)
		}
	}

	return products
}

// applyProductFilters applies the filters to the products slice and returns the filtered slice.
// In a real application, we would apply these filters in the database query (WHERE clause) instead of in-memory.
func (p *productsRepository) applyProductFilters(products []*models.Product, filters *models.ProductFilters) []*models.Product {
	var filteredProducts []*models.Product

	for _, product := range products {
		if filters.Search != "" && !product.Name.Contains(filters.Search) {
			continue
		}

		if filters.Color != "" && !product.Colors.Contains(filters.Color) {
			continue
		}

		if filters.Bestseller != nil && product.Bestseller != *filters.Bestseller {
			continue
		}

		if filters.MinPrice != nil && filters.MaxPrice != nil {
			if product.DiscountedPrice < *filters.MinPrice || product.DiscountedPrice > *filters.MaxPrice {
				continue
			}
		}

		filteredProducts = append(filteredProducts, product)
	}

	return filteredProducts
}

func (p *productsRepository) sortProducts(products []*models.Product, sortBy models.ProductSortBy) []*models.Product {
	switch sortBy {
	case models.SortByPriceAsc:
		slices.SortFunc(products, func(a, b *models.Product) int {
			return cmp.Compare(a.DiscountedPrice, b.DiscountedPrice)
		})

	case models.SortByPriceDesc:
		slices.SortFunc(products, func(a, b *models.Product) int {
			return cmp.Compare(b.DiscountedPrice, a.DiscountedPrice)
		})

	case models.SortByBestseller:
		slices.SortFunc(products, func(a, b *models.Product) int {
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

func (p *productsRepository) GetProducts(_ context.Context, filters *models.ProductFilters, sortBy models.ProductSortBy) (products []*models.Product, err error) {
	metadata, err := p.getProductsMetadata()
	if err != nil {
		return nil, err
	}

	details, err := p.getProductsDetails()
	if err != nil {
		return nil, err
	}

	// This would be a JOIN if we were using a database.
	products = p.joinProductsMetadataAndDetails(metadata, details)

	// Apply filters / WHERE clause in database.
	products = p.applyProductFilters(products, filters)

	// Apply sorting / ORDER BY clause in database.
	products = p.sortProducts(products, sortBy)

	return products, nil
}
