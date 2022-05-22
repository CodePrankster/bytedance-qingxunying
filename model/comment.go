package model

import "gorm.io/gorm"

// 评论结构体

type Comment struct {
	gorm.Model
	Uid        int64  `json:"uid"`                   // 用户id
	Vid        int64  `json:"vid"`                   // 视频id
	Content    string `json:"content,omitempty"`     // 评论内容
	CreateDate string `json:"create_date,omitempty"` // 评论发布日期，格式 mm-dd
}
