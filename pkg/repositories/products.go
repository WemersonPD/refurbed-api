package repositories

import (
	"assignment-backend/pkg/models"
	utils_jsongloader "assignment-backend/pkg/utils/jsongloader"
	utils_numbers "assignment-backend/pkg/utils/numbers"
	"context"
)

const (
	METADATA_FILE_PATH = "data/metadata.json"
	DETAILS_FILE_PATH  = "data/details.json"
)

type ProductsRepository interface {
	// ListProducts reads the metadata and details, merges them and returns the list of products.
	ListProducts(ctx context.Context, filters *models.ProductFilters, sortBy models.ProductSortBy, pagination *models.Pagination) (products []*models.Product, err error)

	// CountProducts return the list of products based on a query.
	CountProducts(_ context.Context, filters *models.ProductFilters) (count int, err error)
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

func (p *productsRepository) findProducts(_ context.Context, filters *models.ProductFilters) (products []*models.Product, err error) {
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
	products = filters.ApplyProductFilters(products)

	return products, err
}

func (p *productsRepository) ListProducts(_ context.Context, filters *models.ProductFilters, sortBy models.ProductSortBy, pagination *models.Pagination) (products []*models.Product, err error) {
	products, err = p.findProducts(context.Background(), filters)
	if err != nil {
		return nil, err
	}

	// Apply sorting / ORDER BY clause in database.
	products = sortBy.SortProducts(products)

	// Apply pagination / LIMIT & OFFSET in database.
	products = pagination.ApplyProductsPagination(products)

	return products, nil
}

func (p *productsRepository) CountProducts(_ context.Context, filters *models.ProductFilters) (count int, err error) {
	products, err := p.findProducts(context.Background(), filters)
	if err != nil {
		return 0, err
	}

	return len(products), err
}
