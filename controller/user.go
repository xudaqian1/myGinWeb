package controller

import (
	"github.com/gin-gonic/gin"
	"myGinWeb/pkg/utils"
	"myGinWeb/service/user_service"
)

type User struct {}
func NewUser() User{
	return User{}
}

type RegisterRequest struct{
	Username string `json:"username" binding:"required"`
	Role string `json:"role" binding:"required,oneof=user admin"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`

}

func (u User) Create(c *gin.Context) {
	var req RegisterRequest
	if err:= c.ShouldBind(&req); err!= nil{
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	salt := utils.GetRandomString(16)
	md5Password := utils.MD5(req.Password + salt)
	userService := user_service.User{
		Username: req.Username,
		Password: md5Password,
		Role: req.Role,
		Email: req.Email,
		Salt: salt,
	}
	if err:=userService.Add();err !=nil{
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "successful",
	})
}

func (u User) GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}
