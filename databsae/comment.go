package databsae

import (
	"time"
)

func CommentCreate(userID int64,videoID int64,text string)(c Comment,err error) {
	//创建实例
	user:=UserInfo{}
	video:=Video{}
	comment:=Comment{
		Content:    text,
		Date: time.Now().Format("01-02"),
		VideoID: videoID,
	}
	//开启事务
	tx:=D.Begin()

	//查询用户
	if res:=tx.Where("id = ?",userID).First(&user);res.Error!=nil{
		tx.Rollback()
		return comment,res.Error
	}else {
		comment.User=user
	}

	//提交评论
	comment.UserID=userID
	if res:=tx.Create(&comment);res.Error!=nil{
		tx.Rollback()
		return comment,res.Error
	}
	//查询视频
	if res:=tx.Select("comment_count").Where("id = ?",videoID).First(&video);res.Error!=nil{
		tx.Rollback()
		return comment,res.Error
	}
	//更新视频评论数
	if res:=tx.Model(&video).Where("id = ?",videoID).Update("comment_count",video.CommentCount+1);res.Error!=nil{
		tx.Rollback()
		return comment,res.Error
	}
	//提交事务
	if res:=tx.Commit();res.Error!=nil{
		return comment,res.Error
	}

	return comment,nil
}

func CommentDelete(videoID string,commentID int64) error  {
	//开启事务
	tx:=D.Begin()
	//视频实例
	video:=Video{}
	//获取视频信息
	if res:=tx.Select("comment_count").Where("id = ?",videoID).First(&video);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新视频评论数
	if res:=tx.Model(&video).Where("id = ?",videoID).Update("comment_count",video.CommentCount-1);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//删除评论
	if res:=tx.Delete(&Comment{},commentID);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//提交事务
	if res:=tx.Commit();res.Error!=nil{
		return res.Error
	}
	return nil
}