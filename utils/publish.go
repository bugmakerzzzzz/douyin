package utils

import (
	"douyin/databsae"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var ip ="http://0.0.0.0:8080/static/"

func Publish(c *gin.Context) {
	token := c.PostForm("token")
	title:=c.PostForm("title")

	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	filename := filepath.Base(data.Filename)
	userId := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", userId, filename)

	saveFile := filepath.Join("./static/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	user:=databsae.UserInfo{}
	databsae.D.Model(&databsae.User{}).Where("id = ?",userId).First(&user)

	video:=databsae.Video{
		AuthorId: user.Id,
		//Author: user,
		PlayUrl: ip+finalName,
		CoverUrl: ip+"2222.png",
		Title:title ,
		FavoriteCount: 0,
		CommentCount: 0,
	}
	res:=databsae.D.Create(&video)
	if res.Error!=nil{
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  res.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}


func PublishList(c *gin.Context) {
	token:=c.Query("token")

	id, exist := usersLoginInfo[token]
	if  !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	videos:=[]databsae.Video{}
	videosInfo:=[]databsae.VideoInfo{}


	res:=databsae.D.Preload("Author").Find(&videos,"author_id = ?",id)
	if res.Error!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
	}
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

	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videosInfo,
	})
}

