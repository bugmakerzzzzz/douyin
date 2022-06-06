package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(c *gin.Context) {
	//获取token
	token := c.Query("token")
	//判断用户权限
	if id, exist := usersLoginInfo[token]; exist {
		//检索用户信息
		userinfo,err:=databsae.SearchUserInfo(id)
		//返回信息
		if err!=nil{
			c.JSON(http.StatusOK, UserInfoResponse{
				Response: Response{StatusCode: 1,StatusMsg:err.Error()},
			})
			return
		}
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 0},
			User:userinfo     ,
		})
		return
	} else {
		c.JSON(http.StatusOK, UserInfoResponse{
			Response: Response{StatusCode: 1, StatusMsg: "User doesn't exist"},
		})
		return
	}
}