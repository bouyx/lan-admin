package handler

import (
	"encoding/json"
	"net/http"
)

func User(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": "logged"})

	}
}
