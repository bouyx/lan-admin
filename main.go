package main

import (
	"fmt"
	"os"

	//"net/http"
	"github.com/bouyx/lan-admin/handler"
	"github.com/bouyx/lan-admin/data"
	//"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	_"github.com/lib/pq"
)

func main() {
	//DB
	data.Start()

	//Vanilla
	// http.HandleFunc("/api/login", handler.Login)
	// port:= os.Getenv("PORT")
	// fmt.Println("listen : "+port)
	// http.ListenAndServe(":"+port, nil)

	//mux 
	// r := mux.NewRouter()
	// r.HandleFunc("/login", handler.Loginmux).Methods("POST")
	// port:= os.Getenv("PORT")
	// fmt.Println("listen : "+port)
	// http.ListenAndServe(":"+port, r)


	router := gin.Default()
	v1 := router.Group("/api/v1")
	v1.Use(gin.Logger())
	v1.GET("/login", handler.Logingin)
	v1.GET("/users", handler.GetAllUsersHandler)
	v1.GET("/users/:id", handler.GetUsersHandler)
	v1.POST("/users", handler.PostUsersHandler)
	v1.PUT("/users/:id", handler.PutUsersHandler)
	port:= os.Getenv("PORT")
	fmt.Println("listen : "+ port)
	router.Run(":"+port)
}
