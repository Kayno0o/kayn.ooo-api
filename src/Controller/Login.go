package controller

import (
	"encoding/json"
	"net/http"

	auth "kayn.ooo/api/src/Auth"
	entity "kayn.ooo/api/src/Entity"
	middleware "kayn.ooo/api/src/Middleware"
)

func Login(mux *http.ServeMux) {
	mux.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		var form entity.UserRegisterForm

		err := json.NewDecoder(r.Body).Decode(&form)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Invalid JSON", "code": "invalid_json", "status": "400"}, 400)
			return
		}

		user, err := auth.Authenticate(form.Email, form.Password)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Invalid email or password", "code": "invalid_credentials", "status": "401"}, 401)
			return
		}

		jwt, err := auth.GenerateToken(user)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
			return
		}

		middleware.WriteJSON(w, jwt, 200)
	})
}
