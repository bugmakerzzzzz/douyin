package databsae

func Publish(id int64,name string,ip string,title string) error {
	//开启事务
	tx := D.Begin()
	//创建用户实例
	user:=UserInfo{}
	//查询用户
	if res:=tx.Where("id = ?",id).First(&user);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//创建video实例
	video:=Video{
		AuthorID: user.ID,
		Author: user,
		PlayUrl: ip+name,
		CoverUrl: ip+"2222.png",
		Title:title ,
		FavoriteCount: 0,
		CommentCount: 0,
		IsFavorite: false,
	}
	//提交到数据库
	if res:=tx.Create(&video);res.Error!=nil{
		tx.Rollback()
		return res.Error
	}
	//提交事务
	return tx.Commit().Error
}