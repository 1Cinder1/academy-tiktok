package router

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/videofeed"
)

type UserRouter struct{}

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/videofeed")
	{
		userRouter.POST("", videofeed.VideoFeed)
		//userRouter.POST("/login", )
	}

	return userRouter
}

func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes {
	userRouter := router.Group("/videofeed")

	return userRouter
}
