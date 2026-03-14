package models

import "strings"

type Name string

func (n Name) ToString() string {
	return string(n)
}

func (n Name) Contains(substring string) bool {
	nameLower := strings.ToLower(n.ToString())
	substringLower := strings.ToLower(substring)

	return strings.Contains(nameLower, substringLower)
}

type Colors []string

func (colors Colors) Contains(color string) bool {
	for _, c := range colors {
		if strings.EqualFold(c, color) {
			return true
		}
	}

	return false
}

type ProductMetadata struct {
	ID        string  `json:"id"`
	Name      Name    `json:"name"`
	BasePrice float64 `json:"base_price"`
	ImageURL  string  `json:"image_url"`
}

type ProductDetail struct {
	ID              string `json:"id"`
	DiscountPercent int    `json:"discount_percent"`
	Bestseller      bool   `json:"bestseller"`
	Colors          Colors `json:"colors"`
	Stock           int    `json:"stock"`
}

type Product struct {
	ID              string  `json:"id"`
	DiscountedPrice float64 `json:"discounted_price"`

	// Metadata fields
	Name      Name    `json:"name"`
	BasePrice float64 `json:"base_price"`
	ImageURL  string  `json:"image_url"`

	// Detail fields
	DiscountPercent int    `json:"discount_percent"`
	Bestseller      bool   `json:"bestseller"`
	Colors          Colors `json:"colors"`
	Stock           int    `json:"stock"`
}
