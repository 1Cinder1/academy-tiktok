package router

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/actionList"
)

type UserRouter struct{}

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/publish")
	{
		userRouter.POST("/list", actionList.GetVideoList)

	}

	return userRouter
}

func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes {
	userRouter := router.Group("/user")
	return userRouter
}
