package models

type ProductMetadata struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	ImageURL  string  `json:"image_url"`
}

type ProductDetail struct {
	ID              string   `json:"id"`
	DiscountPercent int      `json:"discount_percent"`
	Bestseller      bool     `json:"bestseller"`
	Colors          []string `json:"colors"`
	Stock           int      `json:"stock"`
}

type Product struct {
	ID              string  `json:"id"`
	DiscountedPrice float64 `json:"discounted_price"`

	// Metadata fields
	Name      string  `json:"name"`
	BasePrice float64 `json:"base_price"`
	ImageURL  string  `json:"image_url"`

	// Detail fields
	DiscountPercent int      `json:"discount_percent"`
	Bestseller      bool     `json:"bestseller"`
	Colors          []string `json:"colors"`
	Stock           int      `json:"stock"`
}
