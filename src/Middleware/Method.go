package middleware

import (
	"net/http"
)

// Method is a middleware function that only allows requests with specific HTTP methods
func Method(methods ...string) func(next http.Handler) http.Handler {
	allowedMethods := make(map[string]bool)
	for _, method := range methods {
		allowedMethods[method] = true
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !allowedMethods[r.Method] {
				WriteJSON(w, map[string]string{"error": "Method not allowed", "code": "method_not_allowed", "status": "405"}, 405)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
