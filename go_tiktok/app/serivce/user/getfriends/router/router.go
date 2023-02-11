package router

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/global"
	"go_tiktok/app/internal/middleware"
)

func InitRouter() *gin.Engine {

	r := gin.Default()

	r.Use(middleware.ZapLogger(global.Logger), middleware.ZapRecovery(global.Logger, true))

	routerGroup := new(Group)

	publicGroup := r.Group("/api")
	{
		routerGroup.InitUserSignRouter(publicGroup)
	}

	r.Run(":8081")
	global.Logger.Info("initialize routers successfully!")

	return r

}
