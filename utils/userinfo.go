package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func UserInfo(c *gin.Context) {
	//获取token
	token := c.Query("token")
	ids:=c.Query("user_id")
	id,_:=strconv.ParseInt(ids, 10, 64)
	//判断用户权限
	if uid, exist := usersLoginInfo[token]; exist {
		//检索用户信息
		userinfo,err:=databsae.SearchUserInfo(id)
		//返回信息
		if err!=nil{
			c.JSON(http.StatusOK, UserInfoResponse{
				Response: Response{StatusCode: 1,StatusMsg:err.Error()},
			})
			return
		}
		//获取关注数据
		f:=databsae.FollowRelationship{}
		if res:=databsae.D.Where("follower_id = ? and follow_id = ?",uid,id).Find(&f);res.Error!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
			return
		}

		if f.FollowerId==uid{
			userinfo.IsFollow=true
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