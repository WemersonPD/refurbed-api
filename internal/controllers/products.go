package controllers

import (
	"assignment-backend/pkg/models"
	"assignment-backend/pkg/services"
	utils_response "assignment-backend/pkg/utils/response"
	"net/http"
)

type ProductsController interface {
	// GetProducts handles the GET /products endpoint, returning the list of products.
	GetProducts(w http.ResponseWriter, r *http.Request)
}

type productsController struct {
	productService services.ProductService
}

func NewProductsController() ProductsController {
	return &productsController{
		productService: services.NewProductService(),
	}
}

func (c *productsController) GetProducts(w http.ResponseWriter, r *http.Request) {
	filters := models.NewProductFiltersFromRequest(r)
	sort := models.NewProductSortFromRequest(r)
	pagination, err := models.NewPagination(r)
	if err != nil {
		utils_response.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	ctx := r.Context()
	cacheKey := r.URL.RawQuery
	response, err := c.productService.GetProducts(ctx, cacheKey, filters, sort, pagination)
	if err != nil {
		utils_response.Error(w, http.StatusInternalServerError, "Failed to retrieve products")
		return
	}

	paginatedResponse := utils_response.Pagination[[]*models.Product]{
		Limit:  pagination.Limit,
		Offset: pagination.Offset,
		Total:  response.Count,
		Data:   response.Products,
	}

	utils_response.SuccessPaginated(w, paginatedResponse)
}
