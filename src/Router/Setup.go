package router

import (
	"net/http"

	controller "kayn.ooo/api/src/Controller"
	middleware "kayn.ooo/api/src/Middleware"
)

var Mux *http.ServeMux

func InitRoutes() {
	Mux = http.NewServeMux()

	middleware.Api(Mux)

	controller.Register(Mux)
	controller.Login(Mux)
	controller.Welcome(Mux)
	controller.User(Mux)
	controller.Translation(Mux)

	Mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		middleware.WriteJSON(w, map[string]string{"error": "Not found", "code": "not_found", "status": "404"}, 404)
	})
}
