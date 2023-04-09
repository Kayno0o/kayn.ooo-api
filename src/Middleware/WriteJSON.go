package middleware

import (
	"encoding/json"
	"net/http"
)

func WriteJSON(w http.ResponseWriter, data interface{}, status int) {
	if w.Header().Get("Content-Type") == "" {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(status)
		json.NewEncoder(w).Encode(data)
	}
}
