package main

import (
	"douyin/utils"
	"github.com/gin-gonic/gin"
)

func initRouter(r *gin.Engine) {

	// public directory is used to serve static resources
	r.Static("/static", "./public")

	apiRouter := r.Group("/douyin")

	// basic apis

	apiRouter.POST("/user/register/", utils.Register)
	apiRouter.POST("/user/login/", utils.Login)
}