package utils_numbers

func CalculateDiscountedPrice(basePrice float64, discountPercent int) float64 {
	discount := basePrice * (float64(discountPercent) / 100.0)

	return basePrice - discount
}
