package databsae

import (
	"gorm.io/gorm"
	"time"
)

//用户对象
type User struct {
	gorm.Model
	Name  string `gorm:"unique" json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}

//检索容器
type UserInfo struct {
	ID            int64  `gorm:"primaryKey" json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

//视频对象
type Video struct {
	ID            int64  `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt       int64  `gorm:"autoCreateTime"`
	AuthorID      int64
	Author        UserInfo  `json:"author" gorm:"foreignKey:AuthorID"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	Title         string `json:"title,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type FavoriteRelationship struct {
	ID      int64  `gorm:"primaryKey" json:"id,omitempty"`
	UserId  int64
	VideoId int64
}

type Comment struct {
	ID            int64  `gorm:"primaryKey" json:"id,omitempty"`
	CreatedAt     time.Time
	DeletedAt     gorm.DeletedAt `gorm:"index"`
	UserID        int64
	VideoID       int64
	User          UserInfo`json:"user" gorm:"foreignKey:UserID"`
	Date          string  `json:"create_date,omitempty"`
	Content       string  `json:"content,omitempty"`
}

type FollowRelationship struct {
	ID         int64  `gorm:"primaryKey" json:"id,omitempty"`
	FollowId   int64
	FollowerId int64
}