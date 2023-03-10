package router

import (
	"github.com/gin-gonic/gin"
	"go_tiktok/app/serivce/user/register"
)

type UserRouter struct{}

func (r *UserRouter) InitUserSignRouter(router *gin.RouterGroup) gin.IRouter {
	userRouter := router.Group("/user")
	{
		userRouter.POST("/register", register.Register)
		//userRouter.POST("/login", )
	}

	return userRouter
}

func (r *UserRouter) InitUserInfoRouter(router *gin.RouterGroup) gin.IRoutes {
	userRouter := router.Group("/user")

	return userRouter
}
