package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

func Feed(c *gin.Context) {
	timeUnixstr := c.Query("latest_time")
	timeUnix,_:=strconv.ParseInt(timeUnixstr, 10, 64)

	if timeUnixstr==""{
		timeUnix=time.Now().Unix()
	}
	lastestTime:=time.Unix(timeUnix,0).Format("2022-01-01 15:04:05")

	videos:=[]databsae.Video{}
	videosInfo:=[]databsae.VideoInfo{}
	//lastvideo:=databsae.Video{}


	databsae.D.Order("updated_at desc").Preload("Author").Limit(5).Find(&videos,"updated_at <= ?",lastestTime)
	//query.Scan(&videosInfo)
	for _,v:=range videos{
		videosInfo=append(videosInfo,databsae.VideoInfo{
			Id:int64(v.ID),
			Author:databsae.UserInfo{
				Id: v.AuthorId,
				Name: v.Author.Name,
			},
			PlayUrl:v.PlayUrl,
			CoverUrl:v.CoverUrl,
			Title:v.Title,
			FavoriteCount:v.FavoriteCount,
			CommentCount:v.CommentCount,
			IsFavorite:false,
		})
	}
	var newTime int64
	if len(videosInfo)>0{
		newTime=videos[len(videos)-1].UpdatedAt.Unix()
	}else{
		newTime=time.Now().Unix()
	}


	c.JSON(http.StatusOK, FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: videosInfo,
		NextTime:  newTime,
	})
}