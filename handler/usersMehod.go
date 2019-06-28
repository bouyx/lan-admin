package handler

import (
	"github.com/bouyx/lan-admin/data"
)

func getAllUsers()[]data.UserData{

	return data.GetAllUsers()
}

func getUsers(id int)data.UserData{

	return data.GetUsers(id)
}

func createUser(user data.UserData)int{
	return data.CreateUser(user)
}