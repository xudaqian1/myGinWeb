package main

import (
	"fmt"
	"myGinWeb/pkg/setting"
	"myGinWeb/routers"
)

func main() {
	app := routers.SetupRouter()
	if err := app.Run(fmt.Sprintf(":%d", setting.Config.Port)); err != nil {
		fmt.Printf("server startup failed, err:%v\n", err)
	}
}
