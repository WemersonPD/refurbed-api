package utils_response

import (
	"encoding/json"
	"net/http"
)

type ApiResponse[T any] struct {
	OK   bool `json:"ok"`
	Data T    `json:"data"`
}

func toJson[T any](w http.ResponseWriter, status int, data T) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	json.NewEncoder(w).Encode(&ApiResponse[T]{
		OK:   status >= 200 && status < 300,
		Data: data,
	})
}

func Success[T any](w http.ResponseWriter, data T) {
	toJson(w, http.StatusOK, data)
}

func NotFound(w http.ResponseWriter, message string) {
	toJson(w, http.StatusNotFound, map[string]string{"error": message})
}

func Error(w http.ResponseWriter, status int, message string) {
	toJson(w, status, map[string]string{"error": message})
}
