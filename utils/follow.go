package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func RelationAction(c *gin.Context) {
	//获取token，userID
	token := c.Query("token")
	userTo:=c.Query("to_user_id")
	action:=c.Query("action_type")

	//转string为int64
	uto,_:=strconv.ParseInt(userTo, 10, 64)

	//用户鉴权
	uid, exist := usersLoginInfo[token]
	if  !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
		return
	}

	//根据action执行关注
	if action=="1"{
		//关注
		res:=databsae.Follow1(uid,uto)
		if res!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}else{
		//取消关注
		res:=databsae.Follow2(uid,uto)
		if res!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}


func FollowList(c *gin.Context){
	//获取token，id
	token := c.Query("token")
	userID:=c.Query("user_id")

	//用户鉴权
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
		return
	}

	//查询点赞视频列表
	users,err:=databsae.SearchFollowList(userID)

	//返回数据
	if err!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}


func FollowerList(c *gin.Context){
	//获取token，id
	token := c.Query("token")
	userID:=c.Query("user_id")

	//用户鉴权
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
		return
	}

	//查询点赞视频列表
	users,err:=databsae.SearchFollowerList(userID)

	//返回数据
	if err!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, UserListResponse{
		Response: Response{
			StatusCode: 0,
		},
		UserList: users,
	})
}

