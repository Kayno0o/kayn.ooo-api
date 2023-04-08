package controller

import (
	"net/http"

	middleware "kayn.ooo/api/src/Middleware"
)

func Welcome(mux *http.ServeMux) {
	mux.Handle("/welcome",
		middleware.Chain(
			middleware.ValidateToken,
			middleware.Method("GET"),
		)(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Write([]byte("Welcome!"))
		})),
	)
}
