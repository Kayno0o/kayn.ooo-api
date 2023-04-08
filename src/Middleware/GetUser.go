package middleware

import (
	"net/http"

	auth "kayn.ooo/api/src/Auth"
	entity "kayn.ooo/api/src/Entity"
)

// return entity.User or nothing
func GetUser(r *http.Request, w http.ResponseWriter) (*entity.User, error) {
	tokenString := r.Header.Get("Authorization")

	user, err := auth.GetUserFromToken(tokenString)
	if err != nil {
		WriteJSON(w, map[string]string{"error": "Invalid token", "code": "invalid_token", "status": "401"}, 401)
		return nil, err
	}

	return user, nil
}
