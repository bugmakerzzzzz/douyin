package databsae

import (
	"gorm.io/gorm"
)

//用户对象
type User struct {
	gorm.Model
	Name  string `gorm:"unique" json:"name,omitempty"`
	Password string `json:"password,omitempty"`
}
//检索容器
type UserInfo struct {
	Id   int64  `json:"user_id,omitempty"`
	Name  string `json:"name,omitempty"`
}

//视频对象
type Video struct {
	gorm.Model
	AuthorId      int64
	Author        User   `json:"author" gorm:"foreignKey:AuthorId"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	Title         string `json:"title,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
}

type VideoInfo struct {
	Id            int64  `json:"id,omitempty"`
	Author        UserInfo   `json:"author"`
	PlayUrl       string `json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	Title         string `json:"title,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}