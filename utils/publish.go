package utils

import (
	"douyin/databsae"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

var ip ="http://172.19.120.128:8080/static/"

func Publish(c *gin.Context) {
	//获取token，标题
	token := c.PostForm("token")
	title:=c.PostForm("title")

	//验证用户权限
	if _, exist := usersLoginInfo[token]; !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//获取上传文件
	data, err := c.FormFile("data")
	if err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//生成用户名
	filename := filepath.Base(data.Filename)
	userID := usersLoginInfo[token]
	finalName := fmt.Sprintf("%d_%s", userID, filename)

	//保存文件到静态资源
	saveFile := filepath.Join("./static/", finalName)
	if err := c.SaveUploadedFile(data, saveFile); err != nil {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}

	//数据库操作，上传视频
	err=databsae.Publish(userID,finalName,ip,title)

	if err!=nil{
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, Response{
		StatusCode: 0,
		StatusMsg:  finalName + " uploaded successfully",
	})
}


func PublishList(c *gin.Context) {
	//获取token，id
	token:=c.Query("token")
	id:=c.Query("user_id")

	//用户鉴权
	uid, exist := usersLoginInfo[token]
	if  !exist {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
		return
	}

	//创建视频列表实例
	videos:=[]databsae.Video{}

	//创建map容器用于点赞关系
	favorite:=make(map[int64]bool)

	//查询点赞关系表
	r:=[]databsae.FavoriteRelationship{}
	if res:=databsae.D.Where("user_id = ?",id).Find(&r);res.Error!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
	}
	//更新点赞map
	for _,v:=range r{
		favorite[v.VideoId]=true
	}


	follow:=make(map[int64]bool)
	//获取关注数据
	f:=[]databsae.FollowRelationship{}
	if res:=databsae.D.Where("follower_id = ?",uid).Find(&f);res.Error!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
		return
	}
	//更新map
	for _,v:=range f{
		follow[v.FollowId]=true
	}

	//查询视频数据，预加载用户
	res:=databsae.D.Order("created_at desc").Preload("Author").Find(&videos,"author_id = ?",id)
	if res.Error!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
	}

	//更新字段
	for k:=range videos{
		//更新点赞数据
		if _,exist:=favorite[videos[k].ID];exist{
			videos[k].IsFavorite=true
		}//更新关注
		if _,exist:=follow[videos[k].Author.ID];exist{
			videos[k].Author.IsFollow=true
		}

	}

	//返回数据
	c.JSON(http.StatusOK, VideoListResponse{
		Response: Response{
			StatusCode: 0,
		},
		VideoList: videos,
	})
}

