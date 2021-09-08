package controller

import "github.com/gin-gonic/gin"

type User struct {}
func NewUser() User{
	return User{}
}
func (u User) Create(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func (u User) GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}
