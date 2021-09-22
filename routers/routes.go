package routers

import (
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"log"
	"myGinWeb/controller"
	_ "myGinWeb/docs"
	"myGinWeb/middleware"
	"myGinWeb/pkg/setting"
	"myGinWeb/pkg/utils"
)

var user = controller.NewUser()

func SetupRouter() *gin.Engine {
	r := gin.Default()
	gin.SetMode(setting.Config.RunMode)
	limiter, err := utils.NewLimiter(setting.Config.Redis)
	if err != nil {
		log.Fatal(err)
	}
	r.Use(middleware.RateMiddleware(limiter))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
