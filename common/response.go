package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应类型
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}

func OK(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: code,
		StatusMsg:  GetMsg(code),
	})
}
func Error(c *gin.Context, code int32) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: code,
		StatusMsg:  GetMsg(code),
	})
}

func ErrorWithMsg(c *gin.Context, code int32, msg string) {
	c.JSON(http.StatusOK, &Response{
		StatusCode: code,
		StatusMsg:  msg,
	})
}

type Video struct {
	Id            int64  `json:"id,omitempty"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"`
	CoverUrl      string `json:"cover_url,omitempty"`
	FavoriteCount int64  `json:"favorite_count,omitempty"`
	CommentCount  int64  `json:"comment_count,omitempty"`
	IsFavorite    bool   `json:"is_favorite,omitempty"`
}

type User struct {
	Id            int64  `json:"id,omitempty"`
	Name          string `json:"name,omitempty"`
	FollowCount   int64  `json:"follow_count,omitempty"`
	FollowerCount int64  `json:"follower_count,omitempty"`
	IsFollow      bool   `json:"is_follow,omitempty"`
}

type RegAndLogResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	Token      string `json:"token"`       // 用户鉴权token
	UserID     int64  `json:"user_id"`     // 用户id
}

type UserInfoResponse struct {
	StatusCode int64  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string `json:"status_msg"`  // 返回状态描述
	User       User   `json:"user"`        // 用户信息
}
