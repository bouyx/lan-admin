package data

import (
	"os"
	"fmt"
	"database/sql"
)

//remote
// const (
// 	host     = "ec2-54-247-178-166.eu-west-1.compute.amazonaws.com"
// 	port     = "5432"
// 	user     = "bezndebsgvcvwv"
// 	password = "692ef0db5ffae312f5430650ce25468d2c8b890958c390fca22551cddc34a3a2"
// 	dbname   = "d2ln1u49tklonv"
//	ssl		 = "enable"	
// )

//local
const (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "postgres"
	dbname   = "lanManager"
	ssl 	 = "disable"
)

var db *sql.DB
var dberror error

func Start(){
	setEnv("DATABASE_HOST", host)
	setEnv("DATABASE_PORT", port)
	setEnv("DATABASE_USER", user)
	setEnv("DATABASE_PASSWORD", password)
	setEnv("DATABASE_NAME", dbname)
	setEnv("DATABASE_SSL", ssl)

	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
	"password=%s dbname=%s sslmode=%s", os.Getenv("DATABASE_HOST"), os.Getenv("DATABASE_PORT"), os.Getenv("DATABASE_USER"), 
	os.Getenv("DATABASE_PASSWORD"), os.Getenv("DATABASE_NAME"),os.Getenv("DATABASE_SSL"))
	db, dberror = sql.Open("postgres", psqlInfo)
	if dberror != nil {
		panic(dberror)
	}
	dberror = db.Ping()
	if dberror != nil {
		panic(dberror)
	}
	fmt.Println("db connected")
}

func setEnv(name,val string){
	err:= os.Setenv(name, val)
	if err != nil {
		panic(err)
	}
}

func GetAllUsers()(users []UserData){
	rows, err := db.Query("SELECT id,name,email FROM users")
	defer rows.Close()
	for rows.Next() {
	  var data UserData
	  err = rows.Scan(&data.ID,&data.Username,&data.Email)
	  if err != nil {
		// handle this error
		panic(err)
	  }
	  users = append(users,data)
	}
	// get any error encountered during iteration
	err = rows.Err()
	if err != nil {
	  panic(err)
	}
	return
}

func GetUsers(id int)(user UserData){
	row:= db.QueryRow("SELECT id,name,email FROM users WHERE id = $1",id)
	err := row.Scan(&user.ID,&user.Username,&user.Email)
	switch err {
	  case sql.ErrNoRows:
		fmt.Println("No rows were returned!")
		return
	  case nil:
		return
	  default:
		panic(err)
	}
}

func CreateUser(user UserData)(id int){ 
	err := db.QueryRow(`INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id`, user.Username, user.Email).Scan(&id)
	if err != nil {
		panic(err)
	}
	return
}

func UpdateUser(user UserData){ 
	_,err := db.Exec(`UPDATE users SET name = $2, email = $3 WHERE id = $1`, user.ID, user.Username, user.Email)
	if err != nil {
		panic(err)
	}
}

func DeleteUser(id int){ 
	_,err := db.Exec(`DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		panic(err)
	}
}

