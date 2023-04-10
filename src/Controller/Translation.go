package controller

import (
	"net/http"

	entity "kayn.ooo/api/src/Entity"
	middleware "kayn.ooo/api/src/Middleware"
)

func Translation(mux *http.ServeMux) {
	var translations []entity.Translation
	mux.Handle("/translations",
		middleware.Chain(
			middleware.Method("GET"),
			middleware.IsGranted("ROLE_ADMIN"),
		)(GetAllEntitiesHandler(&translations)),
	)

	var translation entity.Translation
	mux.Handle("/translation",
		middleware.Chain(
			middleware.IsGranted("ROLE_ADMIN"),
			middleware.Method("POST", "DELETE", "PUT"),
		)(UpdateEntityHandler(&translation)),
	)
}
