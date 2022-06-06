package databsae

import "gorm.io/gorm/clause"

func Follow1(userId,userTo int64) error{
	//创建关系实例
	f:=FollowRelationship{
		FollowerId: userId,
		FollowId: userTo,
	}
	//开启事务
	tx:=D.Begin()
	//关系点赞关系
	if res:=tx.Create(&f);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//查询用户
	follow:=UserInfo{}
	if res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?",userTo).Select("id","follower_count").Find(&follow);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新数据
	if res:=tx.Model(&follow).Where("id = ?",userTo).Update("follower_count",follow.FollowerCount+1);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//查询用户
	follower:=UserInfo{}
	if res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?",userId).Select("id","follow_count").Find(&follower);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新数据
	if res:=tx.Model(&follower).Where("id = ?",userId).Update("follow_count",follower.FollowCount+1);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//提交事务
	return tx.Commit().Error
}


func Follow2(userId,userTo int64)error {
	//开启事务
	tx:=D.Begin()
	//根据id删除点赞关系
	if res:=tx.Where("follower_id = ? AND follow_id = ?",userId,userTo).Delete(&FollowRelationship{});res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//查询用户
	follow:=UserInfo{}
	if res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?",userTo).Select("id","follower_count").Find(&follow);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新数据
	if res:=tx.Model(&follow).Where("id = ?",userTo).Update("follower_count",follow.FollowerCount-1);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//查询用户
	follower:=UserInfo{}
	if res:=tx.Clauses(clause.Locking{Strength: "UPDATE"}).Where("id = ?",userId).Select("id","follow_count").Find(&follower);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//更新数据
	if res:=tx.Model(&follower).Where("id = ?",userId).Update("follow_count",follower.FollowCount-1);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//提交事务
	return tx.Commit().Error
}

func SearchFollowList(userid string) ([]UserInfo,error) {
	//创建实例
	r:=[]FollowRelationship{}
	users:=[]UserInfo{}

	//查询点赞关系
	if res:=D.Where("follower_id = ?",userid).Find(&r);res.Error!=nil{
		return users,res.Error
	}

	//更新follow切片
	follow:=[]int64{}
	for _,v :=range r{
		follow=append(follow,v.FollowId)
	}

	//查询视频数据，预加载用户
	if res:=D.Find(&users,"id IN ?",follow);res.Error!=nil{
		return users,res.Error
	}

	//更新数据
	for k,_:=range users{
		//更新点赞数据
		users[k].IsFollow=true
	}

	return users,nil
}


func SearchFollowerList(userid string) ([]UserInfo,error) {
	//创建实例
	r:=[]FollowRelationship{}
	users:=[]UserInfo{}
	f:=[]FollowRelationship{}

	//查询关系
	if res:=D.Where("follow_id = ?",userid).Find(&r);res.Error!=nil{
		return users,res.Error
	}
	if res:=D.Where("follower_id = ?",userid).Find(&f);res.Error!=nil{
		return users,res.Error
	}

	//更新follow切片
	follower:=[]int64{}
	for _,v :=range r{
		follower=append(follower,v.FollowerId)
	}
	follow:=make(map[int64]bool)
	for _,v :=range f{
		follow[v.FollowId]=true
	}
	//加载用户
	if res:=D.Find(&users,"id IN ?",follower);res.Error!=nil{
		return users,res.Error
	}

	//更新数据
	for k,_:=range users{
		//更新关注数据
		if _,exist:=follow[users[k].ID];exist{
			users[k].IsFollow=true
		}
	}

	return users,nil
}


