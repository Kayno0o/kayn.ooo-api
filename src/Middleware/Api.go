package middleware

import "net/http"

// make url starts with /api

func Api(mux *http.ServeMux) {
	mux.Handle("/api/", http.StripPrefix("/api", mux))
}
