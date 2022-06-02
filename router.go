package main

import (
	"douyin/utils"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	// public directory is used to serve static resources
	r.Static("/static", "./static")

	apiRouter := r.Group("/douyin")

	// basic apis
	apiRouter.POST("/user/register/", utils.Register)
	apiRouter.POST("/user/login/", utils.Login)
	apiRouter.GET("/user/", utils.UserInfo)
	apiRouter.GET("/feed/", utils.Feed)
	apiRouter.POST("/publish/action/", utils.Publish)
	apiRouter.GET("/publish/list/", utils.PublishList)
}