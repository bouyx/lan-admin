package handler

import (
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
	c.JSON(200, gin.H{
		"message": "suce",
	})
}

func PutUsersHandler(c *gin.Context){
	id := c.Param("id")
	c.JSON(200, gin.H{
		"message": id,
	})
}
