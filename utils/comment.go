package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func CommentAction(c *gin.Context) {
	//获取提交数据
	token := c.Query("token")
	actionType := c.Query("action_type")
	videoID:=c.Query("video_id")
	//用户鉴权
	if userID, exist := usersLoginInfo[token]; exist {
		if actionType == "1" {
			//获取评论内容
			text := c.Query("comment_text")
			vid,_:=strconv.ParseInt(videoID, 10, 64)
			//提交评论
			comment,err:=databsae.CommentCreate(userID,vid,text)
			if err!=nil{
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
				return
			}
			c.JSON(http.StatusOK, CommentActionResponse{Response: Response{StatusCode: 0},
				Comment:comment })
		}else {
			//获取需要删除评论的id
			commentID:=c.Query("comment_id")
			cID,_:=strconv.ParseInt(commentID, 10, 64)
			//删除评论
			err:=databsae.CommentDelete(videoID,cID)
			if err!=nil{
				c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: err.Error()})
				return
			}
			c.JSON(http.StatusOK, Response{StatusCode: 0})
		}
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}


func CommentList(c *gin.Context) {
	//获取信息
	token := c.Query("token")
	videoID:=c.Query("video_id")

	//用户鉴权
	uid, exist := usersLoginInfo[token]

	//创建评论实例
	comments:=[]databsae.Comment{}

	//读取评论列表
	if res:=databsae.D.Order("created_at desc").Preload("User").Select("id","user_id","content","date").Find(&comments,"video_id = ?",videoID);res.Error!=nil{
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: res.Error.Error()})
		return
	}

	if exist{
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

		for k:=range comments{
			//更新关注
			if _,exist:=follow[comments[k].UserID];exist{
				comments[k].User.IsFollow=true
			}

		}
	}


	c.JSON(http.StatusOK, CommentListResponse{
		Response:    Response{StatusCode: 0},
		CommentList: comments,
	})
}