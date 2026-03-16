package utils_numbers_test

import (
	utils_numbers "assignment-backend/pkg/utils/numbers"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateDiscountedPrice(t *testing.T) {
	t.Run("Applying a 25% discount", func(t *testing.T) {
		result := utils_numbers.CalculateDiscountedPrice(100.0, 25)

		assert.Equal(t, 75.0, result)
	})

	t.Run("Applying a 0% discount", func(t *testing.T) {
		result := utils_numbers.CalculateDiscountedPrice(499.99, 0)

		assert.Equal(t, 499.99, result)
	})

	t.Run("Applying a 100% discount", func(t *testing.T) {
		result := utils_numbers.CalculateDiscountedPrice(250.0, 100)

		assert.Equal(t, 0.0, result)
	})

	t.Run("Applying a 50% discount", func(t *testing.T) {
		result := utils_numbers.CalculateDiscountedPrice(610.99, 50)

		assert.Equal(t, 305.495, result)
	})

	t.Run("Applying a discount to a zero price", func(t *testing.T) {
		result := utils_numbers.CalculateDiscountedPrice(0.0, 25)

		assert.Equal(t, 0.0, result)
	})
}
