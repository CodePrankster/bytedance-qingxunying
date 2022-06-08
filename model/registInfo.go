package model

import "gorm.io/gorm"

type RegistInfo struct {
	gorm.Model
	UserName string `json:"user_name,omitempty"` // 用户名
	Password string `json:"password,omitempty"`  // 密码
}
