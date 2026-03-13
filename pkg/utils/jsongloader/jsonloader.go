package utils_jsongloader

import (
	"encoding/json"
	"os"
)

type JSONLoader interface {
	LoadJSONFile(filePath string, expectedType any) error
}

type jsonLoader struct{}

func NewJSONLoader() JSONLoader {
	return &jsonLoader{}
}

// LoadJSONFile reads a JSON file from the given path and unmarshals it into the provided expectedType.
// expectedType should be a pointer to the struct or slice you want to populate with the JSON data.
func (j *jsonLoader) LoadJSONFile(filePath string, expectedType any) (err error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	return json.Unmarshal(data, expectedType)
}
