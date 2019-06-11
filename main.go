package main

import (
	"fmt"
	"os"
	"net/http"
	"github.com/bouyx/lan-admin/handler"
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
	// psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
	// 	"password=%s dbname=%s sslmode=disable",
	// 	host, port, user, password, dbname)
	// db, err := sql.Open("postgres", psqlInfo)
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// fmt.Println("db connected")

	http.HandleFunc("/api/login", handler.Login)

	fmt.Println("listen 8080")
	// http.ListenAndServe(":8080", nil)

	router := gin.New()
	router.Use(gin.Logger())
	router.GET("/login", handler.Logingin)
	port:= os.Getenv("PORT")
	fmt.Println("listen : "+ port)
	router.Run(":"+port)
}
