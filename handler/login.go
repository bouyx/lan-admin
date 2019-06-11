package handler

import (
	// "encoding/json"
	// "net/http"
	// "github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	
)

// func Login(w http.ResponseWriter, r *http.Request) {
// 	if r.Method == "POST" {
// 		w.Header().Set("Content-Type", "application/json")

// 		w.WriteHeader(200)
// 		json.NewEncoder(w).Encode(map[string]string{"message": "logged"})

// 	}
// }

// func Loginmux(w http.ResponseWriter, r *http.Request) {
// 	params := mux.Vars(r)
	
// 	w.Header().Set("Content-Type", "application/json")

// 		w.WriteHeader(200)
// 		json.NewEncoder(w).Encode(map[string]string{"message": params["zboub"]})

// }

func Logingin(c *gin.Context){
	c.JSON(200, gin.H{
		"message": "suce",
	})
}