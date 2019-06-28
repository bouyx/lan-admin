package handler

import (
	"github.com/bouyx/lan-admin/data"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetAllUsersHandler(c *gin.Context){
	c.JSON(200, getAllUsers(),)
}

func GetUsersHandler(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	c.JSON(200, getUsers(id),)
}

func PostUsersHandler(c *gin.Context){
	var user data.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	id:=createUser(user)
	c.JSON(200, getUsers(id),)
}

func PutUsersHandler(c *gin.Context){
	var user data.UserData
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	updateUser(user)
	c.JSON(200, getUsers(user.ID),)
}

func DeleteUsersHandler(c *gin.Context){
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		panic(err)
	}
	deleteUser(id)
	c.Status(200)
}
