package common

import "mime/multipart"

// FavoriteActionRequest 点赞请求参数
type FavoriteActionRequest struct {
	UserId     int64  `json:"user_id"`     // 用户id
	VideoId    int64  `json:"video_id"`    // 视频id
	Token      string `json:"token"`       // 用户鉴权
	ActionType int32  `json:"action_type"` // 点赞类型  1-点赞，2-取消点赞
}

// FavoriteListRequest 点赞请求参数
type FavoriteListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"` // 用户id
	Token  string `json:"token" form:"token"`     // 用户鉴权
}

// PublishActionRequest 发布视频请求参数
type PublishActionRequest struct {
	Data  *multipart.FileHeader `form:"data"`  // 上传的视频
	Token string                `form:"token"` // 用户鉴权
	Title string                `form:"title"` // 视频的标题
}

// PublishListRequest 发布列表请求参数
type PublishListRequest struct {
	Token  string `json:"token"`   // 用户鉴权
	UserId int64  `json:"user_id"` // 用户id
}

// RegistRequest 注册请求参数
type RegAndLogRequest struct {
	Username string `form:"username" ` // 用户名
	Password string `form:"password" ` // 密码
}

// UserInfoRequese 用户信息请求参数
type UserInfoRequese struct {
	Token  string `json:"token"`  // 用户鉴权
	UserId int64  `json:"userId"` // 用户id
}
