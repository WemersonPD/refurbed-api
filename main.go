package main

import (
	"log"
	"net/http"
)

func main() {
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

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		// Placeholder response
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"products":[]}`))
	})

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

