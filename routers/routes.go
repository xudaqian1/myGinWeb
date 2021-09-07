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
	return r
}
