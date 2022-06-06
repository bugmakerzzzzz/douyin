package utils

import (
	"douyin/databsae"
	"github.com/gin-gonic/gin"
	"net/http"
)

var usersLoginInfo=make(map[string]int64)

func Register(c *gin.Context)  {
	//获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	//对密码进行md5加盐处理
	password=MD5_SALT(password)

	//创建用户实例
	var user databsae.User
	user.Name=username
	user.Password=password

	//开启事务，保存到user表
	tx:=databsae.D.Begin()
	if res:=tx.Create(&user);res.Error!=nil{
		tx.Rollback()
	}
	//创建用户信息实例
	userInfo:=databsae.UserInfo{
		ID: int64(user.ID),
		Name: username,
		FollowerCount: 0,
		FollowCount: 0,
	}
	//保存实例到userInfo表
	if res:=tx.Create(&userInfo);res.Error!=nil{
		tx.Rollback()
	}
	//提交事务
	if res:=tx.Commit();res.Error!=nil{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg:"创建用户失败"},
		})
	}else{
		//保存token到内存
		t:=MD5_SALT(username) + password
		usersLoginInfo[t]=userInfo.ID

		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int64(int(user.ID)),
			Token:    t,
		})
	}
}



func Login(c *gin.Context)   {
	//获取用户名密码
	username := c.Query("username")
	password := c.Query("password")
	//获取加密密码
	password=MD5_SALT(password)

	//查询验证密码
	user,err:=databsae.UserLogin(username,password)

	//错误信息
	if err!=""{
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg:err},
		})
		return
	}
	//保存token到内存
	t:=MD5_SALT(username)+password
	usersLoginInfo[t]=int64(user.ID)
	//验证成功响应
	c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 0},
			UserId: int64(user.ID),
			Token:    t,
		})

}

