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
}
