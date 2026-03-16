package main

import (
	"assignment-backend/internal/controllers"
	utils_middlewares "assignment-backend/pkg/utils/middlewares"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	productsController := controllers.NewProductsController(nil)

	mux := http.NewServeMux()

	// Health check endpoint
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"status":"ok"}`))
	})

	// TODO: Implement /products endpoint
	// - Read data/metadata.json and data/details.json
	// - Merge by id
	// - Apply filters from query params (search, color, bestseller, minPrice, maxPrice)
	// - Add caching with 30s TTL

	mux.HandleFunc("/products", productsController.GetProducts)

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", utils_middlewares.CorsMiddleware(utils_middlewares.RecoverErrorMiddleware(mux))); err != nil {
		log.Fatal(err)
	}
}
