package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name     string `json:"name,omitempty"`
	IsFollow bool   `json:"is_follow,omitempty"` // true-已关注，false-未关注
}
