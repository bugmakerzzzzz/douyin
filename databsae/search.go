package databsae

func UserLogin(name,password string)(user User,err string) {
	u:=User{}
	res:=D.Select("password","name","id").Where("name = ?",name).First(&u)
	if res.Error!=nil{
		return u,res.Error.Error()
	}else if u.Password!=password{
		return u,"密码错误"
	}
	return u,""
}

func SearchUserInfo(id int64) (userInfo UserInfo,err string) {
	u:=UserInfo{}
	res:=D.Model(&User{}).First(&u,id)
	if res.Error!=nil{
		return u,res.Error.Error()
	}
	return u,""
}