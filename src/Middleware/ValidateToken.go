package middleware

import (
	"net/http"
	"strings"

	auth "kayn.ooo/api/src/Auth"
)

func ValidateToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")

		if tokenString == "" {
			WriteJSON(w, map[string]string{"error": "Authorization header is required", "code": "authorization_header_required", "status": "401"}, 401)
			return
		}

		token, err := auth.VerifyToken(strings.Replace(tokenString, "Bearer ", "", 1))
		if err != nil {
			WriteJSON(w, map[string]string{"error": "Invalid token", "code": "invalid_token", "status": "401"}, 401)
			return
		}

		if !token.Valid {
			WriteJSON(w, map[string]string{"error": "Invalid token", "code": "invalid_token", "status": "401"}, 401)
			return
		}

		// Pass the request to the next middleware/handler in the chain
		next.ServeHTTP(w, r)
	})
}
