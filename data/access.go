package data

import (
	"fmt"
	"database/sql"
)

const (
	host     = "ec2-54-247-178-166.eu-west-1.compute.amazonaws.com"
	port     = 5432
	user     = "bezndebsgvcvwv"
	password = "692ef0db5ffae312f5430650ce25468d2c8b890958c390fca22551cddc34a3a2"
	dbname   = "d2ln1u49tklonv"
)
var db *sql.DB
var dberror error

func GetAllUsers()(users []UserData){
	rows, err := db.Query("SELECT * FROM users")
	defer rows.Close()
	for rows.Next() {
	  var data UserData
	  err = rows.Scan(data)
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
	row:= db.QueryRow("SELECT * FROM users WHERE id = $1",id)
	err := row.Scan(user)
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

func Start(){
psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
"password=%s dbname=%s", host, port, user, password, dbname)
db, dberror = sql.Open("postgres", psqlInfo)
if dberror != nil {
	panic(dberror)
}
defer db.Close()

dberror = db.Ping()
if dberror != nil {
panic(dberror)
}
fmt.Println("db connected")
}