package databsae

func SearchUser(name,password string)(userInfo UserInfo,err string) {
	u:=UserInfo{}
	res:=D.Model(&User{}).Select("password","name","id").Where("name = ?",name).First(&u)
	if res.Error!=nil{
		return u,"用户不存在"
	}else if u.Password!=password{
		return u,"密码错误"
	}
	return u,""
}