package routers

import (
	"github.com/gin-gonic/gin"
	"myGinWeb/controller"
	"myGinWeb/middleware"
	"myGinWeb/pkg/setting"
)

var user = controller.NewUser()

func SetupRouter() *gin.Engine {
	gin.SetMode(setting.Config.RunMode)
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	r.POST("/users/login", user.Login)
	r.POST("/users/register", user.Create)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(middleware.JWT())
	{
		apiv1.GET("/users/list", user.GetUserList)
	}
	return r
}
