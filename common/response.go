package common

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Response 响应类型
type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
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
	Author        User   `json:"author"`         // 视频作者信息
	CommentCount  int64  `json:"comment_count"`  // 视频的评论总数
	CoverURL      string `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64  `json:"favorite_count"` // 视频的点赞总数
	ID            int64  `json:"id"`             // 视频唯一标识
	IsFavorite    bool   `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string `json:"play_url"`       // 视频播放地址
	Title         string `json:"title"`          // 视频标题
}

type User struct {
	FollowCount   int64  `json:"follow_count"`   // 关注总数
	FollowerCount int64  `json:"follower_count"` // 粉丝总数
	ID            int64  `json:"id"`             // 用户id
	IsFollow      bool   `json:"is_follow"`      // true-已关注，false-未关注
	Name          string `json:"name"`           // 用户名称
}

type RegAndLogResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	Token      *string `json:"token"`       // 用户鉴权token
	UserId     *int64  `json:"user_id"`     // 用户id
}

type UserInfoResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	User       User    `json:"user"`        // 用户信息
}

type PublicListResponse struct {
	StatusCode int64   `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  string  `json:"status_msg"`  // 返回状态描述
	VideoList  []Video `json:"video_list"`  // 用户发布的视频列表
}
type FollowListAndFollowerListResponse struct {
	StatusCode string  `json:"status_code"` // 状态码，0-成功，其他值-失败
	StatusMsg  *string `json:"status_msg"`  // 返回状态描述
	UserList   []User  `json:"user_list"`   // 用户信息列表
}
