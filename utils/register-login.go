package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
)

var usersLoginInfo=make(map[string]int64)

func Register(c *gin.Context)  {
	username := c.Query("username")
	password := c.Query("password")

	password=MD5_SALT(password)

	var user databsae.User
	user.Name=username
	user.Password=password

	res:=databsae.D.Create(&user)

	if res.Error!=nil{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg:"创建用户失败"},
		})
	}else{
		t:=MD5_SALT(username) + password
		usersLoginInfo[t]=int64(user.ID)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int64(int(user.ID)),
			Token:    t,
		})
	}
}

func Login(c *gin.Context)   {
	username := c.Query("username")
	password := c.Query("password")

	password=MD5_SALT(password)

	user,err:=databsae.UserLogin(username,password)

	if err!=""{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg:err},
		})
	}else{
		t:=MD5_SALT(username)+password
		usersLoginInfo[t]=int64(user.ID)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int64(user.ID),
			Token:    t,
		})
	}
}

