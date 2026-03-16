package repositories_test

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/repositories"
	"assignment-backend/pkg/utils/testutil"
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRepo() (productRepository repositories.ProductsRepository, jsonLoader *testutil.MockJSONLoader) {
	mockJsonLoader := &testutil.MockJSONLoader{}

	return repositories.NewProductsRepository(mockJsonLoader), mockJsonLoader
}

func mockLoadFiles(jsonLoader *testutil.MockJSONLoader, productsMetadata []*models.ProductMetadata, productsDetails []*models.ProductDetail) {
	jsonLoader.On("LoadJSONFile", repositories.METADATA_FILE_PATH, mock.Anything).Run(func(args mock.Arguments) {
		prt := args.Get(1).(*[]*models.ProductMetadata)
		*prt = productsMetadata
	}).Return(nil).Once()
	jsonLoader.On("LoadJSONFile", repositories.DETAILS_FILE_PATH, mock.Anything).Run(func(args mock.Arguments) {
		prt := args.Get(1).(*[]*models.ProductDetail)
		*prt = productsDetails
	}).Return(nil).Once()
}

func TestCountProducts(t *testing.T) {
	repo, jsonLoader := setupRepo()

	productsMetadata := testutil.SampleMetadata()
	productsDetails := testutil.SampleDetails()

	t.Run("Success Count Product without filters", func(t *testing.T) {
		mockLoadFiles(jsonLoader, productsMetadata, productsDetails)

		ctx := context.TODO()

		count, err := repo.CountProducts(ctx, &models.ProductFilters{})
		assert.Nil(t, err)
		assert.Equal(t, 3, count, "Expected count to be 3")
	})

	t.Run("Count products with brand filter", func(t *testing.T) {
		mockLoadFiles(jsonLoader, productsMetadata, productsDetails)

		ctx := context.TODO()

		count, err := repo.CountProducts(ctx, &models.ProductFilters{Brands: []string{"apple"}})
		assert.Nil(t, err)
		assert.Equal(t, 2, count, "Expected count to be 2 for apple brand")
	})

	t.Run("Count products with no matches", func(t *testing.T) {
		mockLoadFiles(jsonLoader, productsMetadata, productsDetails)

		ctx := context.TODO()

		count, err := repo.CountProducts(ctx, &models.ProductFilters{Search: "nonexistent"})
		assert.Nil(t, err)
		assert.Equal(t, 0, count, "Expected count to be 0")
	})

	t.Run("Count products returns error on metadata failure", func(t *testing.T) {
		jsonLoader.On("LoadJSONFile", repositories.METADATA_FILE_PATH, mock.Anything).Return(fmt.Errorf("failed to read metadata.json")).Once()

		ctx := context.TODO()

		count, err := repo.CountProducts(ctx, &models.ProductFilters{})
		assert.Error(t, err)
		assert.Equal(t, 0, count)
	})

	t.Run("Count products returns error on details failure", func(t *testing.T) {
		jsonLoader.On("LoadJSONFile", repositories.METADATA_FILE_PATH, mock.Anything).Run(func(args mock.Arguments) {
			prt := args.Get(1).(*[]*models.ProductMetadata)
			*prt = productsMetadata
		}).Return(nil).Once()
		jsonLoader.On("LoadJSONFile", repositories.DETAILS_FILE_PATH, mock.Anything).Return(fmt.Errorf("failed to read details.json")).Once()

		ctx := context.TODO()

		count, err := repo.CountProducts(ctx, &models.ProductFilters{})
		assert.Error(t, err)
		assert.Equal(t, 0, count)
	})
}

func TestListProducts(t *testing.T) {
	repo, jsonLoader := setupRepo()

	productsMetadata := testutil.SampleMetadata()
	productsDetails := testutil.SampleDetails()

	t.Run("List all products without filters", func(t *testing.T) {
		mockLoadFiles(jsonLoader, productsMetadata, productsDetails)

		ctx := context.TODO()
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		products, err := repo.ListProducts(ctx, &models.ProductFilters{}, models.SortByBestseller, pagination)
		assert.Nil(t, err)
		assert.Len(t, products, 3)
	})

	t.Run("List products with search filter", func(t *testing.T) {
		mockLoadFiles(jsonLoader, productsMetadata, productsDetails)

		ctx := context.TODO()
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		products, err := repo.ListProducts(ctx, &models.ProductFilters{Search: "iphone"}, models.SortByBestseller, pagination)
		assert.Nil(t, err)
		assert.Len(t, products, 1)
		assert.Equal(t, "p1", products[0].ID)
	})

	t.Run("List products with pagination", func(t *testing.T) {
		mockLoadFiles(jsonLoader, productsMetadata, productsDetails)

		ctx := context.TODO()
		pagination := &models.Pagination{Limit: 2, Offset: 0}

		products, err := repo.ListProducts(ctx, &models.ProductFilters{}, models.SortByPriceAsc, pagination)
		assert.Nil(t, err)
		assert.Len(t, products, 2)
	})

	t.Run("List products returns error on metadata failure", func(t *testing.T) {
		jsonLoader.On("LoadJSONFile", repositories.METADATA_FILE_PATH, mock.Anything).Return(fmt.Errorf("failed to read metadata.json")).Once()

		ctx := context.TODO()
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		products, err := repo.ListProducts(ctx, &models.ProductFilters{}, models.SortByBestseller, pagination)
		assert.Error(t, err)
		assert.Nil(t, products)
	})

	t.Run("List products returns error on details failure", func(t *testing.T) {
		jsonLoader.On("LoadJSONFile", repositories.METADATA_FILE_PATH, mock.Anything).Run(func(args mock.Arguments) {
			prt := args.Get(1).(*[]*models.ProductMetadata)
			*prt = productsMetadata
		}).Return(nil).Once()
		jsonLoader.On("LoadJSONFile", repositories.DETAILS_FILE_PATH, mock.Anything).Return(fmt.Errorf("failed to read details.json")).Once()

		ctx := context.TODO()
		pagination := &models.Pagination{Limit: 10, Offset: 0}

		products, err := repo.ListProducts(ctx, &models.ProductFilters{}, models.SortByBestseller, pagination)
		assert.Error(t, err)
		assert.Nil(t, products)
	})
}
