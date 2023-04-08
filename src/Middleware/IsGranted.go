package middleware

import (
	"net/http"

	auth "kayn.ooo/api/src/Auth"
)

func IsGranted(role string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")

			user, err := auth.GetUserFromToken(tokenString)
			if err != nil {
				WriteJSON(w, map[string]string{"error": "Invalid token", "code": "invalid_token", "status": "401"}, 401)
				return
			}

			if !user.HasRole(role) {
				WriteJSON(w, map[string]string{"error": "Forbidden", "code": "forbidden", "status": "403"}, 403)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
