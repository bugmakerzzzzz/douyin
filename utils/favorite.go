package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func FavoriteAction(c *gin.Context) {
	//获取token，userID，videoID
	token := c.Query("token")
	videoId:=c.Query("video_id")
	action:=c.Query("action_type")

	//转string为int64
	vid,_:=strconv.ParseInt(videoId, 10, 64)

	//用户鉴权
	uid, exist := usersLoginInfo[token]
	if  !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
		return
	}

	//根据action执行点赞或取消点赞
	if action=="1"{
		//点赞
		res:=databsae.Favorite1(uid,vid)
		if res!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}else{
		//取消点赞
		res:=databsae.Favorite2(uid,vid)
		if res!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error()})
			return
		}
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	}
}


func FavoriteList(c *gin.Context){
	//获取token，id
	token := c.Query("token")
	userID:=c.Query("user_id")

	//用户鉴权
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't login"})
		return
	}

	//查询点赞视频列表
	videos,err:=databsae.SearchFavoriteList(userID)

	//返回数据
	if err!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
		return
	}
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}