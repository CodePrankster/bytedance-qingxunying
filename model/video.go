package model

import "gorm.io/gorm"

type Video struct {
	gorm.Model
	Uid        uint   `json:"uid"`                                // 用户id
	PlayUrl    string `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	CoverUrl   string `json:"cover_url,omitempty"`                // 视频封面地址
	Title      string `json:"title,omitempty"`                    // 视频标题
	IsFavorite bool   `json:"is_favorite,omitempty"`              // true-已点赞，false-未点赞
}
