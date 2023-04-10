package controller

import (
	"encoding/json"
	"net/http"

	middleware "kayn.ooo/api/src/Middleware"
	repository "kayn.ooo/api/src/Repository"
)

func GetAllEntitiesHandler(entities interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := repository.FindAll(entities)
		if err != nil {
			middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
			return
		}

		middleware.WriteJSON(w, entities, 200)
	}
}

func UpdateEntityHandler(entity interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			err := json.NewDecoder(r.Body).Decode(entity)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Invalid JSON", "code": "invalid_json", "status": "400"}, 400)
				return
			}

			err = repository.Create(entity)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
				return
			}

			middleware.WriteJSON(w, entity, 200)
			return
		}

		if r.Method == "DELETE" {
			err := json.NewDecoder(r.Body).Decode(entity)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Invalid JSON", "code": "invalid_json", "status": "400"}, 400)
				return
			}

			err = repository.Delete(entity)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
				return
			}

			middleware.WriteJSON(w, map[string]string{"message": "Entity deleted"}, 200)
			return
		}

		if r.Method == "PUT" {
			err := json.NewDecoder(r.Body).Decode(entity)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Invalid JSON", "code": "invalid_json", "status": "400"}, 400)
				return
			}

			err = repository.Update(entity)
			if err != nil {
				middleware.WriteJSON(w, map[string]string{"error": "Internal server error", "code": "internal_server_error", "status": "500"}, 500)
				return
			}

			middleware.WriteJSON(w, map[string]string{"message": "Entity created"}, 200)
			return
		}
	}
}
