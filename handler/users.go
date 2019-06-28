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
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}
