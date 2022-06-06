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
	User databsae.UserInfo `json:"user,omitempty"`
}

type FeedResponse struct {
	Response
	VideoList []databsae.Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

type VideoListResponse struct {
	Response
	VideoList []databsae.Video `json:"video_list"`
}

type CommentListResponse struct {
	Response
	CommentList []databsae.Comment `json:"comment_list,omitempty"`
}

type CommentActionResponse struct {
	Response
	Comment databsae.Comment `json:"comment,omitempty"`
}

type UserListResponse struct {
	Response
	UserList []databsae.UserInfo `json:"user_list"`
}