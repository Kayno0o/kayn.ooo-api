package controller

import (
	"net/http"

	middleware "kayn.ooo/api/src/Middleware"
	repository "kayn.ooo/api/src/Repository"
)

func User(mux *http.ServeMux) {
	mux.Handle("/users",
		middleware.Chain(
			middleware.Method("GET"),
			middleware.IsGranted("ROLE_ADMIN"),
		)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			users, err := repository.GetUsers()
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
				return
			}

			middleware.WriteJSON(w, users, 200)
		})),
	)

	// get connected user
	mux.Handle("/current_user",
		middleware.Chain(
			middleware.Method("GET"),
		)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			user, err := middleware.GetUser(r, w)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
				return
			}

			middleware.WriteJSON(w, user, 200)
		}),
		),
	)
}
