package routers

import (
	"github.com/gin-gonic/gin"
	"myGinWeb/controller"
	"myGinWeb/pkg/setting"
)

func SetupRouter() *gin.Engine {
	gin.SetMode(setting.Config.RunMode)
	r := gin.Default()
	r.GET("/ping", controller.Ping)
	apiv1 := r.Group("/api/v1")
	{
		user := controller.NewUser()
		apiv1.POST("/users/register",user.Create)
		apiv1.GET("/users/list",user.GetUserList)
	}
	return r
}
