package utils

import "douyin/databsae"

//响应结构体
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

//用户登录与注册
type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

//用户信息
type UserInfoResponse struct {
	Response
	User databsae.UserInfo
}

type FeedResponse struct {
	Response
	VideoList []databsae.VideoInfo `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []databsae.VideoInfo `json:"video_list"`
}