package databsae

func UserLogin(name,password string)(user User,err string) {
	u:=User{}
	res:=D.Where("name = ?",name).First(&u)
	if res.Error!=nil{
		return u,res.Error.Error()
	}else if u.Password!=password{
		return u,"密码错误"
	}
	return u,""
}


//根据id获取用户信息
func SearchUserInfo(id int64) (userInfo UserInfo,err error) {
	u:=UserInfo{}
	res:=D.First(&u,id)
	if res.Error!=nil{
		return u,res.Error
	}
	return u,nil
}