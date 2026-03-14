package utils_middlewares

import (
	utils_response "assignment-backend/pkg/utils/response"
	"net/http"
)

func RecoverErrorMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			err := recover()

			if err == nil {
				return
			}

			if e, ok := err.(error); ok {
				utils_response.Error(w, http.StatusInternalServerError, e.Error())

				return
			}

			if e, ok := err.(string); ok {
				utils_response.Error(w, http.StatusInternalServerError, e)

				return
			}

			utils_response.Error(w, http.StatusInternalServerError, "An unexpected error occurred")
		}()

		next.ServeHTTP(w, r)
	})
}
