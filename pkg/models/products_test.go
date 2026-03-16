package models_test

import (
	"assignment-backend/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestName_ToString(t *testing.T) {
	t.Run("Returns string representation", func(t *testing.T) {
		name := models.Name("iPhone 15")

		assert.Equal(t, "iPhone 15", name.ToString())
	})
}

func TestName_Contains(t *testing.T) {
	t.Run("Exact match", func(t *testing.T) {
		name := models.Name("iPhone 15")

		assert.True(t, name.Contains("iPhone 15"))
	})

	t.Run("Partial match", func(t *testing.T) {
		name := models.Name("iPhone 15")

		assert.True(t, name.Contains("iPhone"))
	})

	t.Run("Case insensitive match", func(t *testing.T) {
		name := models.Name("iPhone 15")

		assert.True(t, name.Contains("iphone"))
	})

	t.Run("No match", func(t *testing.T) {
		name := models.Name("iPhone 15")

		assert.False(t, name.Contains("Galaxy"))
	})
}

func TestColors_Contains(t *testing.T) {
	colors := models.Colors{"blue", "red", "green"}

	t.Run("Color exists", func(t *testing.T) {
		assert.True(t, colors.Contains("blue"))
	})

	t.Run("Color does not exist", func(t *testing.T) {
		assert.False(t, colors.Contains("black"))
	})

	t.Run("Case insensitive match", func(t *testing.T) {
		assert.True(t, colors.Contains("BLUE"))
	})
}

func TestCategory_Contains(t *testing.T) {
	category := models.Category("smartphones")

	t.Run("Category in list", func(t *testing.T) {
		assert.True(t, category.Contains([]string{"smartphones", "tablets"}))
	})

	t.Run("Category not in list", func(t *testing.T) {
		assert.False(t, category.Contains([]string{"laptops", "accessories"}))
	})

	t.Run("Case insensitive match", func(t *testing.T) {
		assert.True(t, category.Contains([]string{"SMARTPHONES"}))
	})
}

func TestBrand_Contains(t *testing.T) {
	brand := models.Brand("apple")

	t.Run("Brand in list", func(t *testing.T) {
		assert.True(t, brand.Contains([]string{"apple", "samsung"}))
	})

	t.Run("Brand not in list", func(t *testing.T) {
		assert.False(t, brand.Contains([]string{"google", "xiaomi"}))
	})

	t.Run("Case insensitive match", func(t *testing.T) {
		assert.True(t, brand.Contains([]string{"APPLE"}))
	})
}

func TestCondition_Contains(t *testing.T) {
	condition := models.Condition("refurbished")

	t.Run("Condition in list", func(t *testing.T) {
		assert.True(t, condition.Contains([]string{"new", "refurbished"}))
	})

	t.Run("Condition not in list", func(t *testing.T) {
		assert.False(t, condition.Contains([]string{"new", "used"}))
	})

	t.Run("Case insensitive match", func(t *testing.T) {
		assert.True(t, condition.Contains([]string{"REFURBISHED"}))
	})
}
