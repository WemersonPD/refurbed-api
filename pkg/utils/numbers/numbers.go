package utils_numbers

func CalculateDiscountedPrice(basePrice float64, discountPercent int) float64 {
	if discountPercent == 0 {
		return basePrice
	}

	discount := basePrice * (float64(discountPercent) / 100.0)

	return basePrice - discount
}
