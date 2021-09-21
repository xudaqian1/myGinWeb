package controller

import (
	"github.com/gin-gonic/gin"
	"myGinWeb/pkg/utils"
	"myGinWeb/service/user_service"
)

type User struct{}

func NewUser() User {
	return User{}
}

type RegisterRequest struct {
	Username string `json:"username" binding:"required"`
	Role     string `json:"role" binding:"required,oneof=user admin"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Create @Summary 注册
// @Description 通过接口进行注册
// @Tags 用户
// @Accept json
// @Param username body string true "用户名"
// @Param role body string true "权限" Enums("user", "admin")
// @Param email body string true "邮箱"
// @Param password body string true "密码"
// @Success 200 {object} gin.H "{"code":200,"data":{},"msg":"ok"}"
// @Failure 400 {object} gin.H "{"msg": "create failure"}"
// @Router /users/login [post]
func (u User) Create(c *gin.Context) {
	var req RegisterRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}

	salt := utils.GetRandomString(16)
	userService := user_service.User{
		Username: req.Username,
		Password: req.Password,
		Role:     req.Role,
		Email:    req.Email,
		Salt:     salt,
	}
	data := make(map[string]interface{})
	if err := userService.Add(); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	token, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	data["token"] = token
	c.JSON(200, gin.H{
		"data": data,
	})
}

type LoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBind(&req); err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	userService := user_service.User{
		Username: req.Username,
		Password: req.Password,
	}
	isExist, err := userService.CheckAuth()
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	if !isExist {
		c.JSON(400, gin.H{
			"msg": "密码错误",
		})
		return
	}
	token, err := utils.GenerateToken(req.Username)
	if err != nil {
		c.JSON(400, gin.H{
			"msg": err.Error(),
		})
		return
	}
	data := make(map[string]interface{})
	data["token"] = token
	c.JSON(200, gin.H{
		"data": data,
		"msg":  "login successful",
	})
}
func (u User) GetUserList(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "list",
	})
}
