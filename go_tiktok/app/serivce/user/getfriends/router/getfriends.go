package router

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/getfriends"
)

type UserRouter struct{}

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/relation")
	{
		userRouter.POST("/friend/list", getfans.GetFriends)

	}

	return userRouter
}

func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes {
	userRouter := router.Group("/user")

	return userRouter
}
