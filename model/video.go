package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Uid           int64  `json:"uid"`                                // 用户id
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`                // 视频封面地址
	Title         string `json:"title,omitempty"`                    // 视频标题
	FavoriteCount int64  `json:"favorite_count,omitempty"`           // 视频的点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`            // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty"`              // true-已点赞，false-未点赞
}
