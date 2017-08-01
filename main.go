package main

import (
	"fmt"
	"testserver/routers"

	"github.com/littletwolee/commons"

	"github.com/gin-gonic/gin"
)

func main() {
	gin.SetMode(commons.Config.GetString("system.runmode"))
	router := gin.New()
	router.Use(gin.Logger())
	routers.Router.InitRouters(router)
	err := router.Run(fmt.Sprintf("%s:%d", commons.Config.GetString("system.host"), commons.Config.GetInt("system.port")))
	if err != nil {
		commons.GetLogger().LogPanic(err)
	}
}
