package testutil

import "github.com/stretchr/testify/mock"

// MockJSONLoader is a testify mock for the JSONLoader interface.
type MockJSONLoader struct {
	mock.Mock
}

func (m *MockJSONLoader) LoadJSONFile(filePath string, expectedType any) error {
	args := m.Called(filePath, expectedType)
	return args.Error(0)
}
