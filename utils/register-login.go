package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Register(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	var user databsae.User
	user.Name=username
	user.Password=password

	res:=databsae.D.Create(&user)

	if res.Error!=nil{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg:"创建用户失败"},
		})
	}else{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int64(int(user.ID)),
			Token:    username + password,
		})
	}
}

func Login(c *gin.Context)   {
	username := c.Query("username")
	password := c.Query("password")

	userinfo,err:=databsae.SearchUser(username,password)

	if err!=""{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg:err},
		})
	}else{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int64(userinfo.Id),
			Token:    username+password,
		})
	}
}

