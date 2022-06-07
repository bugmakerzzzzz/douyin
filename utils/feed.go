package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	//获取时间戳，token，将时间转为int64
	timeUnixstr := c.Query("latest_time")
	token:=c.Query("token")
	timeUnix,_:=strconv.ParseInt(timeUnixstr, 10, 64)
	//判断时间为空，填充为当前时间
	if timeUnixstr=="0" || timeUnixstr==""{
		timeUnix=time.Now().Unix()
	}

	//创建数据容器
	videos:=[]databsae.Video{}

	//用于判断是否favorite
	favorite:=make(map[int64]bool)
	//判断是否follow
	follow:=make(map[int64]bool)
	//获取点赞关系表，填充map容器
	if userID,exist:=usersLoginInfo[token];exist{
		r:=[]databsae.FavoriteRelationship{}
		//获取关系表
		if res:=databsae.D.Where("user_id = ?",userID).Find(&r);res.Error!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
			return
		}
		//添加map内容
		for _,v:=range r{
			favorite[v.VideoId]=true
		}
		//获取关注数据
		f:=[]databsae.FollowRelationship{}
		if res:=databsae.D.Where("follower_id = ?",userID).Find(&f);res.Error!=nil{
			c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
			return
		}
		//更新map
		for _,v:=range f{
			follow[v.FollowId]=true
		}
	}

	//查询videos，预加载user
	databsae.D.Order("created_at desc").Preload("Author").Limit(5).Find(&videos,"created_at <= ?",timeUnix)

	//更新字段
	for k,_:=range videos{
		//更新点赞数据
		if _,exist:=favorite[videos[k].ID];exist{
			videos[k].IsFavorite=true
		}
		//更新关注
		if _,exist:=follow[videos[k].Author.ID];exist{
			videos[k].Author.IsFollow=true
		}
	}

	//更新时间
	var newTime =time.Now().Unix()
	if len(videos)>0{
		newTime=videos[len(videos)-1].CreatedAt
	}

	//返回数据
	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videos,
		NextTime:  newTime,
	})
}