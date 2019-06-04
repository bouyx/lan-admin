package handler

import (
	"encoding/json"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "logged"})

	}
}
