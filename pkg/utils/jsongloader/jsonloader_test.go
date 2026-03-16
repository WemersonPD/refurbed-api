package utils_jsongloader_test

import (
	utils_jsongloader "assignment-backend/pkg/utils/jsongloader"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadJSONFile(t *testing.T) {
	t.Run("Successfully loading a JSON file", func(t *testing.T) {
		file, _ := os.CreateTemp("", "*.json")
		defer os.Remove(file.Name())
		file.WriteString(`[{"id": "p1", "name": "iPhone"}]`)
		file.Close()

		var result []struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		}

		loader := utils_jsongloader.NewJSONLoader()
		err := loader.LoadJSONFile(file.Name(), &result)

		assert.NoError(t, err)
		assert.Len(t, result, 1)
		assert.Equal(t, "p1", result[0].ID)
		assert.Equal(t, "iPhone", result[0].Name)
	})

	t.Run("File not found", func(t *testing.T) {
		loader := utils_jsongloader.NewJSONLoader()
		var result []struct{}

		err := loader.LoadJSONFile("nonexistent.json", &result)

		assert.Error(t, err)
	})

	t.Run("Invalid JSON content", func(t *testing.T) {
		file, _ := os.CreateTemp("", "*.json")
		defer os.Remove(file.Name())
		file.WriteString(`{invalid json}`)
		file.Close()

		loader := utils_jsongloader.NewJSONLoader()
		var result map[string]string

		err := loader.LoadJSONFile(file.Name(), &result)

		assert.Error(t, err)
	})
}
