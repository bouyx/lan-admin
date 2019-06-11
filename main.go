package main

import (
	"fmt"
	"os"
	//"net/http"
	"github.com/bouyx/lan-admin/handler"
	//"github.com/gorilla/mux"
	"github.com/gin-gonic/gin"
	//_ "github.com/lib/pq"
)

// const (
// 	host     = "localhost"
// 	port     = 5432
// 	user     = "postgres"
// 	password = "postgres"
// 	dbname   = "lanManager"
// )

func main() {
	//DB

	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// fmt.Println("db connected")

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


	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/login", handler.Logingin)
	port:= os.Getenv("PORT")
	fmt.Println("listen : "+ port)
	router.Run(":"+port)
}
