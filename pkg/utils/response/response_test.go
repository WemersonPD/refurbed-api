package utils_response_test

import (
	utils_response "assignment-backend/pkg/utils/response"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	t.Run("Returns 200 with ok true and data", func(t *testing.T) {
		w := httptest.NewRecorder()

		utils_response.Success(w, map[string]string{"name": "iPhone 15"})

		assert.Equal(t, http.StatusOK, w.Code)
		assert.Equal(t, "application/json", w.Header().Get("Content-Type"))

		var body map[string]interface{}
		json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, true, body["ok"])
		assert.Equal(t, "iPhone 15", body["data"].(map[string]interface{})["name"])
	})
}

func TestSuccessPaginated(t *testing.T) {
	t.Run("Returns 200 with paginated data", func(t *testing.T) {
		w := httptest.NewRecorder()

		utils_response.SuccessPaginated(w, utils_response.Pagination[[]string]{
			Limit:  6,
			Offset: 0,
			Total:  15,
			Data:   []string{"p1", "p2"},
		})

		assert.Equal(t, http.StatusOK, w.Code)

		var body map[string]interface{}
		json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, true, body["ok"])

		data := body["data"].(map[string]interface{})
		assert.Equal(t, float64(6), data["limit"])
		assert.Equal(t, float64(0), data["offset"])
		assert.Equal(t, float64(15), data["total"])
		assert.Len(t, data["data"].([]interface{}), 2)
	})
}

func TestNotFound(t *testing.T) {
	t.Run("Returns 404 with error message", func(t *testing.T) {
		w := httptest.NewRecorder()

		utils_response.NotFound(w, "product not found")

		assert.Equal(t, http.StatusNotFound, w.Code)

		var body map[string]interface{}
		json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, false, body["ok"])
		assert.Equal(t, "product not found", body["data"].(map[string]interface{})["error"])
	})
}

func TestError(t *testing.T) {
	t.Run("Returns 400 with error message", func(t *testing.T) {
		w := httptest.NewRecorder()

		utils_response.Error(w, http.StatusBadRequest, "invalid request")

		assert.Equal(t, http.StatusBadRequest, w.Code)

		var body map[string]interface{}
		json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, false, body["ok"])
		assert.Equal(t, "invalid request", body["data"].(map[string]interface{})["error"])
	})

	t.Run("Returns 500 with error message", func(t *testing.T) {
		w := httptest.NewRecorder()

		utils_response.Error(w, http.StatusInternalServerError, "internal error")

		assert.Equal(t, http.StatusInternalServerError, w.Code)

		var body map[string]interface{}
		json.NewDecoder(w.Body).Decode(&body)

		assert.Equal(t, false, body["ok"])
		assert.Equal(t, "internal error", body["data"].(map[string]interface{})["error"])
	})
}
