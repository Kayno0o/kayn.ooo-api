package controller

import (
	"encoding/json"
	"net/http"
	"regexp"

	"golang.org/x/crypto/bcrypt"
	auth "kayn.ooo/api/src/Auth"
	entity "kayn.ooo/api/src/Entity"
	middleware "kayn.ooo/api/src/Middleware"
	repository "kayn.ooo/api/src/Repository"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		var form entity.UserRegisterForm

		err := json.NewDecoder(r.Body).Decode(&form)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Invalid JSON", "code": "invalid_json", "status": "400"}, 400)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
			return
		}

		form.Password = string(hashedPassword)

		if form.Email == "" || form.Password == "" {
			middleware.WriteJSON(w, map[string]string{"error": "Email and password are required", "code": "invalid_credentials", "status": "400"}, 400)
			return
		}

		re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		if !re.MatchString(form.Email) {
			middleware.WriteJSON(w, map[string]string{"error": "Email is invalid", "code": "invalid_credentials", "status": "400"}, 400)
			return
		}

		var existingUser entity.User
		userByEmail := repository.DB.Where("email = ?", form.Email).First(&existingUser)
		if userByEmail.RowsAffected > 0 {
			middleware.WriteJSON(w, map[string]string{"error": "Email is already taken", "code": "invalid_credentials", "status": "400"}, 400)
			return
		}

		user := entity.User{
			Email:    form.Email,
			Password: form.Password,
		}

		result := repository.DB.Create(&user)
		if result.Error != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
			return
		}

		jwt, err := auth.GenerateToken(&user)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
			return
		}

		middleware.WriteJSON(w, jwt, 200)
	})
}
