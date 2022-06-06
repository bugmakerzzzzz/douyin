package databsae

import "gorm.io/gorm/clause"

func Favorite1(userId,videoId int64) error{
	//创建关系实例
	f:=FavoriteRelationship{
		UserId: userId,
		VideoId: videoId,
	}
	//开启事务
	tx:=D.Begin()
	//关系点赞关系
	if res:=tx.Create(&f);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//查询视频
	video:=Video{}
	if res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?",videoId).Find(&video);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新评论点赞数
	video.FavoriteCount++
	if res:=tx.Save(&video);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//提交事务
	return tx.Commit().Error
}


func Favorite2(userId,videoId int64) error {
	//开启事务
	tx:=D.Begin()
	//根据id删除点赞关系
	if res:=tx.Where("user_id = ? AND video_id = ?",userId,videoId).Delete(&FavoriteRelationship{});res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//查询视频信息
	video:=Video{}
	if res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?",videoId).Find(&video);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新点赞数
	video.FavoriteCount--
	if res:=tx.Save(&video);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//提交事务
	return tx.Commit().Error
}


func SearchFavoriteList(userid string) ([]Video,error) {
	//创建实例
	r:=[]FavoriteRelationship{}
	videos:=[]Video{}

	//查询点赞关系
	if res:=D.Where("user_id = ?",userid).Find(&r);res.Error!=nil{
		return videos,res.Error
	}

	//更新favorite切片
	favorite:=[]int64{}
	for _,v :=range r{
		favorite=append(favorite,v.VideoId)
	}

	//查询视频数据，预加载用户
	if res:=D.Preload("Author").Find(&videos,"id IN ?",favorite);res.Error!=nil{
		return videos,res.Error
	}

	//更新数据
	for k,_:=range videos{
		//更新点赞数据
		videos[k].IsFavorite=true
	}

	return videos,nil
}