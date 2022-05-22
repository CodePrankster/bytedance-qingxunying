package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount int64  `json:"follower_count,omitempty"` // 粉丝总数
	IsFollow      bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
}
