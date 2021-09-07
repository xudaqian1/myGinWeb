package main

import (
	"fmt"
	"myGinWeb/models"
	"myGinWeb/pkg/setting"
	"myGinWeb/routers"
)

func main() {
	err:=models.InitDb(setting.Config.Database)
	if err !=nil{
		fmt.Printf("init mysql failed, err:%v\n", err)
		return
	}
	defer models.CloseDB()
	app := routers.SetupRouter()
	if err := app.Run(fmt.Sprintf(":%d", setting.Config.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
