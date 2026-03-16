package testutil

import "assignment-backend/pkg/models"

func BoolPtr(b bool) *bool          { return &b }
func Float64Ptr(f float64) *float64 { return &f }

func SampleProducts() []*models.Product {
	return []*models.Product{
		{
			ID:              "p1",
			Name:            "iPhone 15",
			BasePrice:       999.99,
			DiscountedPrice: 749.99,
			Bestseller:      true,
			Colors:          models.Colors{"blue", "red"},
			Category:        "smartphones",
			Brand:           "apple",
			Condition:       "refurbished",
		},
		{
			ID:              "p2",
			Name:            "Galaxy S24",
			BasePrice:       899.99,
			DiscountedPrice: 629.99,
			Bestseller:      false,
			Colors:          models.Colors{"black", "green"},
			Category:        "smartphones",
			Brand:           "samsung",
			Condition:       "new",
		},
		{
			ID:              "p3",
			Name:            "MacBook Pro",
			BasePrice:       1499.99,
			DiscountedPrice: 1199.99,
			Bestseller:      true,
			Colors:          models.Colors{"silver", "black"},
			Category:        "laptops",
			Brand:           "apple",
			Condition:       "used",
		},
	}
}

func SampleMetadata() []*models.ProductMetadata {
	return []*models.ProductMetadata{
		{ID: "p1", Name: "iPhone 15", BasePrice: 1000.0, ImageURL: "iphone.jpg"},
		{ID: "p2", Name: "Galaxy S24", BasePrice: 800.0, ImageURL: "galaxy.jpg"},
		{ID: "p3", Name: "MacBook Pro", BasePrice: 1500.0, ImageURL: "macbook.jpg"},
	}
}

func SampleDetails() []*models.ProductDetail {
	return []*models.ProductDetail{
		{ID: "p1", DiscountPercent: 25, Bestseller: true, Colors: models.Colors{"blue", "red"}, Stock: 10, Category: "smartphones", Brand: "apple", Condition: "refurbished"},
		{ID: "p2", DiscountPercent: 10, Bestseller: false, Colors: models.Colors{"black"}, Stock: 5, Category: "smartphones", Brand: "samsung", Condition: "new"},
		{ID: "p3", DiscountPercent: 20, Bestseller: true, Colors: models.Colors{"silver"}, Stock: 3, Category: "laptops", Brand: "apple", Condition: "used"},
	}
}
