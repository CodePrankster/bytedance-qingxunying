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

// RegAndLogRequest 注册请求参数
type RegAndLogRequest struct {
	Username string `form:"username" ` // 用户名
	Password string `form:"password" ` // 密码
}

// UserInfoRequese 用户信息请求参数
type UserInfoRequese struct {
	Token  string `json:"token"`  // 用户鉴权
	UserId int64  `json:"userId"` // 用户id
}

// RelationActionRequest 关系操作参数
type RelationActionRequest struct {
	UserId     int64  `json:"user_id"`     //用户id
	Token      string `json:"token"`       // 用户鉴权
	ToUserId   int64  `json:"to_user_id"`  // 对方用户id
	ActionType int32  `json:"action_type"` // 1-关注 2-取消关注
}

// RelationFollowListRequest 关系操作参数
type RelationFollowListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"` // 用户id
	Token  string `json:"token" form:"token"`     // 用户鉴权
}

// RelationFollowerListRequest 关系操作参数
type RelationFollowerListRequest struct {
	UserId int64  `json:"user_id" form:"user_id"` // 用户id
	Token  string `json:"token" form:"token"`     // 用户鉴权
}

// CommentRequest 发布评论参数
type CommentRequest struct {
	Token       string `json:"token"`
	VideoId     int64  `json:"video_id"`
	ActionType  int32  `json:"action_type"`            // 1-发布评论，2-删除评论
	CommentText string `json:"comment_text,omitempty"` //用户填写的评论内容，在action_type=1的时候使用
	CommentId   int64  `json:"comment_id,omitempty"`   // 要删除的评论id，在action_type=2的时候使用
}

// CommentListRequest 评论列表参数
type CommentListRequest struct {
	Token   string `json:"token" form:"token"`
	VideoId int64  `json:"video_id" form:"video_id"`
}
