package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	token := c.Query("token")

	if id, exist := usersLoginInfo[token]; exist {
		userinfo,err:=databsae.SearchUserInfo(id)
		if err!=""{
			c.JSON(http.StatusOK, UserInfoResponse{
				Response: Response{StatusCode: 1,StatusMsg:err},
			})
		}
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:userinfo     ,
		})
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
	}
}