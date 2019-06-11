package handler

import (
	"encoding/json"
	"net/http"
	"github.com/gorilla/mux"
	
)

// func Login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		w.Header().Set("Content-Type", "application/json")

// 		w.WriteHeader(200)
// 		json.NewEncoder(w).Encode(map[string]string{"message": "logged"})

// 	}
// }

func Loginmux(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	
	w.Header().Set("Content-Type", "application/json")

		w.WriteHeader(200)
		json.NewEncoder(w).Encode(map[string]string{"message": params["zboub"]})

}

